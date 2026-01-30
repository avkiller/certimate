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

func (c *WafClient) ListHost(request *model.ListHostRequest) (*model.ListHostResponse, error) {
	requestDef := hwwaf.GenReqDefForListHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostResponse), nil
	}
}

func (c *WafClient) ListPremiumHost(request *model.ListPremiumHostRequest) (*model.ListPremiumHostResponse, error) {
	requestDef := hwwaf.GenReqDefForListPremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPremiumHostResponse), nil
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

func (c *WafClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := hwwaf.GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

func (c *WafClient) UpdateHost(request *model.UpdateHostRequest) (*model.UpdateHostResponse, error) {
	requestDef := hwwaf.GenReqDefForUpdateHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostResponse), nil
	}
}

func (c *WafClient) UpdatePremiumHost(request *model.UpdatePremiumHostRequest) (*model.UpdatePremiumHostResponse, error) {
	requestDef := hwwaf.GenReqDefForUpdatePremiumHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePremiumHostResponse), nil
	}
}
