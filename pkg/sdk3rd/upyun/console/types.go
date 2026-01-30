package console

import (
	"encoding/json"
)

type sdkResponse interface {
	GetData() *sdkResponseBaseData
}

type sdkResponseBase struct {
	Data *sdkResponseBaseData `json:"data,omitempty"`
}

func (r *sdkResponseBase) GetData() *sdkResponseBaseData {
	return r.Data
}

var _ sdkResponse = (*sdkResponseBase)(nil)

type sdkResponseBaseData struct {
	ErrorCode json.Number `json:"error_code,omitempty"`
	Message   string      `json:"message,omitempty"`
}

func (r *sdkResponseBaseData) GetErrorCode() int {
	if r.ErrorCode.String() == "" {
		return 0
	}

	errcode, err := r.ErrorCode.Int64()
	if err != nil {
		return -1
	}

	return int(errcode)
}

func (r *sdkResponseBaseData) GetMessage() string {
	return r.Message
}
