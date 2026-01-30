package dnscom

import (
	"encoding/json"
)

type sdkResponse interface {
	GetCode() int
	GetMessage() string
}

type sdkResponseBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r *sdkResponseBase) GetCode() int {
	return r.Code
}

func (r *sdkResponseBase) GetMessage() string {
	return r.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)

type DomainRecord struct {
	GroupID        json.Number `json:"groupID"`
	DomainID       json.Number `json:"domainsID"`
	Domain         string      `json:"domains"`
	State          int32       `json:"state"`
	UserLockState  int32       `json:"userLock"`
	AdminLockState int32       `json:"adminLock"`
	HealthState    int32       `json:"healthState"`
	ViewType       string      `json:"view_type"`
}

type DNSRecord struct {
	DomainID json.Number `json:"domainID"`
	RecordID json.Number `json:"recordID"`
	ViewID   json.Number `json:"viewID"`
	Record   string      `json:"record"`
	Type     string      `json:"type"`
	Host     string      `json:"host"`
	Value    string      `json:"value"`
	TTL      int32       `json:"ttl"`
	MX       int32       `json:"mx"`
	State    int32       `json:"state"`
	Remark   string      `json:"remark"`
}
