package v2

import (
	"context"
	"net/http"
)

type WebsiteSSLUploadRequest struct {
	SSLID           int64  `json:"sslID"`
	Type            string `json:"type"`
	Certificate     string `json:"certificate"`
	CertificatePath string `json:"certificatePath"`
	PrivateKey      string `json:"privateKey"`
	PrivateKeyPath  string `json:"privateKeyPath"`
	Description     string `json:"description"`
}

type WebsiteSSLUploadResponse struct {
	sdkResponseBase
}

func (c *Client) WebsiteSSLUpload(req *WebsiteSSLUploadRequest) (*WebsiteSSLUploadResponse, error) {
	return c.WebsiteSSLUploadWithContext(context.Background(), req)
}

func (c *Client) WebsiteSSLUploadWithContext(ctx context.Context, req *WebsiteSSLUploadRequest) (*WebsiteSSLUploadResponse, error) {
	httpreq, err := c.newRequest(http.MethodPost, "/websites/ssl/upload")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &WebsiteSSLUploadResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
