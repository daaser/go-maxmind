package request

import (
	"bytes"
	"encoding/json"
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
