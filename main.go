package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/daaser/go-maxmind/pkg/request"
	"go.uber.org/zap"
)

type JSONReader struct {
	*bytes.Reader
}

func NewJSONReader(v interface{}) *JSONReader {
	jr := new(JSONReader)
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	jr.Reader = bytes.NewReader(buf)
	return jr
}

func newLoggerInfallible() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}

func pprint(v interface{}) {
	body, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body)
}

func main() {
	ctx := context.Background()
	mmc := request.NewClient(ctx)
	resp, err := mmc.PostScore(&request.PostScoreJSONRequestBody{})
	if err != nil {
		panic(err)
	}
	pprint(resp)
}
