// Package geoipupdate provides a library for using MaxMind's GeoIP Update
// service.
package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type MaxMindClient struct {
	ctx  context.Context
	rt   http.RoundTripper
	base *url.URL
}

type MaxMindRequest struct {
	Path   string
	Method string
	Body   interface{}
}

// NewClient creates an *http.Client for use in updating.
func NewClient(ctx context.Context) *MaxMindClient {
	url := &url.URL{
		Scheme: "https",
		Host:   "minfraud.maxmind.com",
		Path:   "/minfraud/v2.0/",
	}
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       5 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	return &MaxMindClient{
		rt:   transport,
		ctx:  ctx,
		base: url,
	}
}

func (mm MaxMindClient) PostScore(body *PostScoreJSONRequestBody) (*PostScoreJSONBody, error) {
	mmr := &MaxMindRequest{
		Path:   "score",
		Method: http.MethodPost,
		Body:   body,
	}
	buf, err := mm.do(mmr)
	if err != nil {
		return nil, err
	}

	resp := new(PostScoreJSONBody)
	if err := json.Unmarshal(buf, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func dumpRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", dump)
}

func dumpResponse(res *http.Response) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", dump)
}

func (mm MaxMindClient) do(mmr *MaxMindRequest) ([]byte, error) {
	var buf []byte
	url, err := mm.base.Parse(mmr.Path)
	if err != nil {
		return nil, err
	}
	jr := NewJSONReader(mmr.Body)

	if err := retryWithBackoff(func() error {
		req, err := http.NewRequest(mmr.Method, url.String(), jr)
		if err != nil {
			return errors.Wrap(err, "error creating HTTP request")
		}

		res, err := mm.rt.RoundTrip(req)
		if err != nil {
			return errors.Wrap(err, "error performing HTTP request")
		}
		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Fatalf("error closing response body: %+v", errors.Wrap(err, "closing body"))
			}
		}()

		buf, err = ioutil.ReadAll(io.LimitReader(res.Body, 1<<20))
		if err != nil {
			return errors.Wrap(err, "error reading response body")
		} else if res.StatusCode != http.StatusOK {
			return errors.Errorf("unexpected HTTP status code: %s", res.Status)
		} else if len(buf) == 0 {
			return errors.New("response body is empty")
		} else if bytes.Count(buf, []byte("\n")) > 0 || bytes.Count(buf, []byte("\x00")) > 0 {
			return errors.New("invalid characters in filename")
		}

		return nil
	}, 5*time.Second); err != nil {
		return nil, err
	}
	return buf, nil
}
