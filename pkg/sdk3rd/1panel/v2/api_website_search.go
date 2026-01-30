package v2

import (
	"context"
	"net/http"
)

type WebsiteSearchRequest struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Order    string `json:"order"`
	OrderBy  string `json:"orderBy"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
}

type WebsiteSearchResponse struct {
	sdkResponseBase

	Data *struct {
		Items []*struct {
			ID            int64  `json:"id"`
			Alias         string `json:"alias"`
			PrimaryDomain string `json:"primaryDomain"`
			Protocol      string `json:"protocol"`
			Type          string `json:"type"`
			Status        string `json:"status"`
			SitePath      string `json:"sitePath"`
			Remark        string `json:"remark"`
			SSLStatus     string `json:"sslStatus"`
			SSLExpireDate string `json:"sslExpireDate"`
			UpdatedAt     string `json:"updatedAt"`
			CreatedAt     string `json:"createdAt"`
		} `json:"items"`
		Total int32 `json:"total"`
	} `json:"data,omitempty"`
}

func (c *Client) WebsiteSearch(req *WebsiteSearchRequest) (*WebsiteSearchResponse, error) {
	return c.WebsiteSearchWithContext(context.Background(), req)
}

func (c *Client) WebsiteSearchWithContext(ctx context.Context, req *WebsiteSearchRequest) (*WebsiteSearchResponse, error) {
	httpreq, err := c.newRequest(http.MethodPost, "/websites/search")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &WebsiteSearchResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
