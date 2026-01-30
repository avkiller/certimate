package baishan

import (
	"encoding/json"
)

type sdkResponse interface {
	GetCode() int
	GetMessage() string
}

type sdkResponseBase struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (r *sdkResponseBase) GetCode() int {
	if r.Code == nil {
		return 0
	}

	return *r.Code
}

func (r *sdkResponseBase) GetMessage() string {
	if r.Message == nil {
		return ""
	}

	return *r.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)

type DomainRecord struct {
	Id         string `json:"id"`
	Domain     string `json:"domain"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Cname      string `json:"cname"`
	Area       string `json:"area"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type DomainCertificate struct {
	CertId         json.Number `json:"cert_id"`
	Name           string      `json:"name"`
	CertStartTime  string      `json:"cert_start_time"`
	CertExpireTime string      `json:"cert_expire_time"`
}

type DomainConfig struct {
	Https *DomainConfigHttps `json:"https"`
}

type DomainConfigHttps struct {
	CertId      json.Number `json:"cert_id"`
	ForceHttps  *string     `json:"force_https,omitempty"`
	EnableHttp2 *string     `json:"http2,omitempty"`
	EnableOcsp  *string     `json:"ocsp,omitempty"`
}
