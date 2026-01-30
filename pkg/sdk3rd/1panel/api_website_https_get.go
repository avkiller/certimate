package onepanel

import (
	"context"
	"fmt"
	"net/http"
)

type WebsiteHttpsGetResponse struct {
	sdkResponseBase

	Data *struct {
		Enable       bool     `json:"enable"`
		WebsiteSSLID int64    `json:"websiteSSLId"`
		HttpConfig   string   `json:"httpConfig"`
		SSLProtocol  []string `json:"SSLProtocol"`
		Algorithm    string   `json:"algorithm"`
		Hsts         bool     `json:"hsts"`
	} `json:"data,omitempty"`
}

func (c *Client) WebsiteHttpsGet(websiteId int64) (*WebsiteHttpsGetResponse, error) {
	return c.WebsiteHttpsGetWithContext(context.Background(), websiteId)
}

func (c *Client) WebsiteHttpsGetWithContext(ctx context.Context, websiteId int64) (*WebsiteHttpsGetResponse, error) {
	if websiteId == 0 {
		return nil, fmt.Errorf("sdkerr: unset websiteId")
	}

	httpreq, err := c.newRequest(http.MethodGet, fmt.Sprintf("/websites/%d/https", websiteId))
	if err != nil {
		return nil, err
	} else {
		httpreq.SetContext(ctx)
	}

	result := &WebsiteHttpsGetResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
