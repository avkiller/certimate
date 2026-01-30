package onepanel

import (
	"context"
	"net/http"
)

type SettingsSSLUpdateRequest struct {
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	SSLType     string `json:"sslType"`
	SSL         string `json:"ssl"`
	SSLID       int64  `json:"sslID"`
	AutoRestart string `json:"autoRestart"`
}

type SettingsSSLUpdateResponse struct {
	sdkResponseBase
}

func (c *Client) SettingsSSLUpdate(req *SettingsSSLUpdateRequest) (*SettingsSSLUpdateResponse, error) {
	return c.SettingsSSLUpdateWithContext(context.Background(), req)
}

func (c *Client) SettingsSSLUpdateWithContext(ctx context.Context, req *SettingsSSLUpdateRequest) (*SettingsSSLUpdateResponse, error) {
	httpreq, err := c.newRequest(http.MethodPost, "/settings/ssl/update")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &SettingsSSLUpdateResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
