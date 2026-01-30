package internal

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	hwelb "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

// This is a partial copy of https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/master/services/elb/v3/elb_client.go
// to lightweight the vendor packages in the built binary.
type ElbClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewElbClient(hcClient *httpclient.HcHttpClient) *ElbClient {
	return &ElbClient{HcClient: hcClient}
}

func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := hwelb.GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := hwelb.GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := hwelb.GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}
