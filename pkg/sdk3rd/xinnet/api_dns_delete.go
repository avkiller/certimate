package xinnet

import (
	"context"
)

type DnsDeleteRequest struct {
	DomainName *string `json:"domainName,omitempty"`
	RecordId   *int64  `json:"recordId,omitempty"`
}

type DnsDeleteResponse struct {
	sdkResponseBase
}

func (c *Client) DnsDelete(req *DnsDeleteRequest) (*DnsDeleteResponse, error) {
	return c.DnsDeleteWithContext(context.Background(), req)
}

func (c *Client) DnsDeleteWithContext(ctx context.Context, req *DnsDeleteRequest) (*DnsDeleteResponse, error) {
	httpreq, err := c.newRequest("/dns/delete/")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &DnsDeleteResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
