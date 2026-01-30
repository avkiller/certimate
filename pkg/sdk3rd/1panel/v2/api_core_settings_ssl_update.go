package v2

import (
	"context"
	"net/http"
)

type CoreSettingsSSLUpdateRequest struct {
	Cert        string `json:"cert"`
	Key         string `json:"key"`
	SSLType     string `json:"sslType"`
	SSL         string `json:"ssl"`
	SSLID       int64  `json:"sslID"`
	AutoRestart string `json:"autoRestart"`
}

type CoreSettingsSSLUpdateResponse struct {
	sdkResponseBase
}

func (c *Client) CoreSettingsSSLUpdate(req *CoreSettingsSSLUpdateRequest) (*CoreSettingsSSLUpdateResponse, error) {
	return c.CoreSettingsSSLUpdateWithContext(context.Background(), req)
}

func (c *Client) CoreSettingsSSLUpdateWithContext(ctx context.Context, req *CoreSettingsSSLUpdateRequest) (*CoreSettingsSSLUpdateResponse, error) {
	httpreq, err := c.newRequest(http.MethodPost, "/core/settings/ssl/update")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &CoreSettingsSSLUpdateResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
