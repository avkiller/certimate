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

func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := hwelb.GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := hwelb.GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := hwelb.GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := hwelb.GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

func (c *ElbClient) ShowLoadBalancer(request *model.ShowLoadBalancerRequest) (*model.ShowLoadBalancerResponse, error) {
	requestDef := hwelb.GenReqDefForShowLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerResponse), nil
	}
}

func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := hwelb.GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}
