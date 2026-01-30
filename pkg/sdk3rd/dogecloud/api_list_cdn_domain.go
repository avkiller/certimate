package dogecloud

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListCdnDomainResponse struct {
	sdkResponseBase

	Data *struct {
		Domains []*struct {
			Id          int64           `json:"id"`
			Name        string          `json:"name"`
			Cname       string          `json:"cname"`
			ServiceType string          `json:"service_type"`
			Status      string          `json:"status"`
			Source      json.RawMessage `json:"source"`
			CreateTime  string          `json:"ctime"`
			CertId      int64           `json:"cert_id"`
		} `json:"domains"`
	} `json:"data,omitempty"`
}

func (c *Client) ListCdnDomain() (*ListCdnDomainResponse, error) {
	return c.ListCdnDomainWithContext(context.Background())
}

func (c *Client) ListCdnDomainWithContext(ctx context.Context) (*ListCdnDomainResponse, error) {
	httpreq, err := c.newRequest(http.MethodGet, "/cdn/domain/list.json")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetContext(ctx)
	}

	result := &ListCdnDomainResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
