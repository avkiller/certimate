package xinnet

import (
	"context"
)

type DnsCreateRequest struct {
	DomainName *string `json:"domainName,omitempty"`
	RecordName *string `json:"recordName,omitempty"`
	Type       *string `json:"type,omitempty"`
	Value      *string `json:"value,omitempty"`
	Line       *string `json:"line,omitempty"`
	Ttl        *int32  `json:"ttl,omitempty"`
	Mx         *int32  `json:"mx,omitempty"`
	Status     *int32  `json:"status,omitempty"`
}

type DnsCreateResponse struct {
	sdkResponseBase
	Data *int64 `json:"data,omitempty"`
}

func (c *Client) DnsCreate(req *DnsCreateRequest) (*DnsCreateResponse, error) {
	return c.DnsCreateWithContext(context.Background(), req)
}

func (c *Client) DnsCreateWithContext(ctx context.Context, req *DnsCreateRequest) (*DnsCreateResponse, error) {
	httpreq, err := c.newRequest("/dns/create/")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &DnsCreateResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
