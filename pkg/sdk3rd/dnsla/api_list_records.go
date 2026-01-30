package dnsla

import (
	"context"
	"net/http"

	qs "github.com/google/go-querystring/query"
)

type ListRecordsRequest struct {
	DomainId  *string `json:"domainId,omitempty"  url:"domainId,omitempty"`
	GroupId   *string `json:"groupId,omitempty"   url:"groupId,omitempty"`
	LineId    *string `json:"lineId,omitempty"    url:"lineId,omitempty"`
	Type      *int32  `json:"type,omitempty"      url:"type,omitempty"`
	Host      *string `json:"host,omitempty"      url:"host,omitempty"`
	Data      *string `json:"data,omitempty"      url:"data,omitempty"`
	PageIndex *int32  `json:"pageIndex,omitempty" url:"pageIndex,omitempty"`
	PageSize  *int32  `json:"pageSize,omitempty"  url:"pageSize,omitempty"`
}

type ListRecordsResponse struct {
	sdkResponseBase
	Data *struct {
		Total   int32        `json:"total"`
		Results []*DnsRecord `json:"results"`
	} `json:"data,omitempty"`
}

func (c *Client) ListRecords(req *ListRecordsRequest) (*ListRecordsResponse, error) {
	return c.ListRecordsWithContext(context.Background(), req)
}

func (c *Client) ListRecordsWithContext(ctx context.Context, req *ListRecordsRequest) (*ListRecordsResponse, error) {
	httpreq, err := c.newRequest(http.MethodGet, "/recordList")
	if err != nil {
		return nil, err
	} else {
		values, err := qs.Values(req)
		if err != nil {
			return nil, err
		}

		httpreq.SetQueryParamsFromValues(values)
		httpreq.SetContext(ctx)
	}

	result := &ListRecordsResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
