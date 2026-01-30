package ucdn

import (
	ucloudcdn "github.com/ucloud/ucloud-sdk-go/services/ucdn"
)

type GetUcdnDomainConfigRequest = ucloudcdn.GetUcdnDomainConfigRequest

type GetUcdnDomainConfigResponse = ucloudcdn.GetUcdnDomainConfigResponse

func (c *UCDNClient) NewGetUcdnDomainConfigRequest() *GetUcdnDomainConfigRequest {
	req := &GetUcdnDomainConfigRequest{}

	c.Client.SetupRequest(req)

	req.SetRetryable(true)
	return req
}

func (c *UCDNClient) GetUcdnDomainConfig(req *GetUcdnDomainConfigRequest) (*GetUcdnDomainConfigResponse, error) {
	var err error
	var res GetUcdnDomainConfigResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUcdnDomainConfig", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
