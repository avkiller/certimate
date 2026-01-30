package ucdn

import (
	ucloudcdn "github.com/ucloud/ucloud-sdk-go/services/ucdn"
)

type UpdateUcdnDomainHttpsConfigV2Request = ucloudcdn.UpdateUcdnDomainHttpsConfigV2Request

type UpdateUcdnDomainHttpsConfigV2Response = ucloudcdn.UpdateUcdnDomainHttpsConfigV2Response

func (c *UCDNClient) NewUpdateUcdnDomainHttpsConfigV2Request() *UpdateUcdnDomainHttpsConfigV2Request {
	req := &UpdateUcdnDomainHttpsConfigV2Request{}

	c.Client.SetupRequest(req)

	req.SetRetryable(true)
	return req
}

func (c *UCDNClient) UpdateUcdnDomainHttpsConfigV2(req *UpdateUcdnDomainHttpsConfigV2Request) (*UpdateUcdnDomainHttpsConfigV2Response, error) {
	var err error
	var res UpdateUcdnDomainHttpsConfigV2Response

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateUcdnDomainHttpsConfigV2", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
