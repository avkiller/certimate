package baishan

import (
	"context"
	"net/http"
)

type UploadDomainCertificateRequest struct {
	CertificateId *string `json:"cert_id,omitempty"`
	Certificate   *string `json:"certificate,omitempty"`
	Key           *string `json:"key,omitempty"`
	Name          *string `json:"name,omitempty"`
}

type UploadDomainCertificateResponse struct {
	sdkResponseBase

	Data *DomainCertificate `json:"data,omitempty"`
}

func (c *Client) UploadDomainCertificate(req *UploadDomainCertificateRequest) (*UploadDomainCertificateResponse, error) {
	return c.UploadDomainCertificateWithContext(context.Background(), req)
}

func (c *Client) UploadDomainCertificateWithContext(ctx context.Context, req *UploadDomainCertificateRequest) (*UploadDomainCertificateResponse, error) {
	httpreq, err := c.newRequest(http.MethodPost, "/v2/domain/certificate")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &UploadDomainCertificateResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
