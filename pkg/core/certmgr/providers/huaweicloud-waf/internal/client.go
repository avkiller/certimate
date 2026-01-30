package internal

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	hwwaf "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/waf/v1/model"
)

// This is a partial copy of https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/master/services/waf/v1/waf_client.go
// to lightweight the vendor packages in the built binary.
type WafClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewWafClient(hcClient *httpclient.HcHttpClient) *WafClient {
	return &WafClient{HcClient: hcClient}
}

func (c *WafClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := hwwaf.GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

func (c *WafClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := hwwaf.GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

func (c *WafClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := hwwaf.GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}
