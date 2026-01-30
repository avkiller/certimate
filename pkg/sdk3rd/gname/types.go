package gname

import "encoding/json"

type sdkResponse interface {
	GetCode() int
	GetMessage() string
}

type sdkResponseBase struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (r *sdkResponseBase) GetCode() int {
	return r.Code
}

func (r *sdkResponseBase) GetMessage() string {
	return r.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)

type DomainResolutionRecordord struct {
	ID          json.Number `json:"id"`
	ZoneName    string      `json:"ym"`
	RecordType  string      `json:"lx"`
	RecordName  string      `json:"zjt"`
	RecordValue string      `json:"jxz"`
	MX          int32       `json:"mx"`
}
