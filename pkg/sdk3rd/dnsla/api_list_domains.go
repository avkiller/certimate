package dnsla

import (
	"context"
	"net/http"

	qs "github.com/google/go-querystring/query"
)

type ListDomainsRequest struct {
	GroupId   *string `json:"groupId,omitempty"   url:"groupId,omitempty"`
	PageIndex *int32  `json:"pageIndex,omitempty" url:"pageIndex,omitempty"`
	PageSize  *int32  `json:"pageSize,omitempty"  url:"pageSize,omitempty"`
}

type ListDomainsResponse struct {
	sdkResponseBase
	Data *struct {
		Total   int32           `json:"total"`
		Results []*DomainRecord `json:"results"`
	} `json:"data,omitempty"`
}

func (c *Client) ListDomains(req *ListDomainsRequest) (*ListDomainsResponse, error) {
	return c.ListDomainsWithContext(context.Background(), req)
}

func (c *Client) ListDomainsWithContext(ctx context.Context, req *ListDomainsRequest) (*ListDomainsResponse, error) {
	httpreq, err := c.newRequest(http.MethodGet, "/domainList")
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

	result := &ListDomainsResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
