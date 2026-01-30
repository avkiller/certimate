package cdnfly

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type sdkResponse interface {
	GetCode() string
	GetMessage() string
}

type sdkResponseBase struct {
	Code    json.RawMessage `json:"code"`
	Message string          `json:"msg"`
}

func (r *sdkResponseBase) GetCode() string {
	if r.Code == nil {
		return ""
	}

	decoder := json.NewDecoder(bytes.NewReader(r.Code))
	token, err := decoder.Token()
	if err != nil {
		return ""
	}

	switch t := token.(type) {
	case string:
		return t
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case json.Number:
		return t.String()
	default:
		return ""
	}
}

func (r *sdkResponseBase) GetMessage() string {
	return r.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)
