package internal

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcwaf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf/v20180125"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/waf/v20180125/client.go
// to lightweight the vendor packages in the built binary.
type WafClient struct {
	common.Client
}

func NewWafClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *WafClient, err error) {
	client = &WafClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *WafClient) DescribeDomainDetailsSaas(request *tcwaf.DescribeDomainDetailsSaasRequest) (response *tcwaf.DescribeDomainDetailsSaasResponse, err error) {
	return c.DescribeDomainDetailsSaasWithContext(context.Background(), request)
}

func (c *WafClient) DescribeDomainDetailsSaasWithContext(ctx context.Context, request *tcwaf.DescribeDomainDetailsSaasRequest) (response *tcwaf.DescribeDomainDetailsSaasResponse, err error) {
	if request == nil {
		request = tcwaf.NewDescribeDomainDetailsSaasRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "waf", tcwaf.APIVersion, "DescribeDomainDetailsSaas")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeDomainDetailsSaas require credential")
	}

	request.SetContext(ctx)
	response = tcwaf.NewDescribeDomainDetailsSaasResponse()
	err = c.Send(request, response)
	return
}

func (c *WafClient) ModifySpartaProtection(request *tcwaf.ModifySpartaProtectionRequest) (response *tcwaf.ModifySpartaProtectionResponse, err error) {
	return c.ModifySpartaProtectionWithContext(context.Background(), request)
}

func (c *WafClient) ModifySpartaProtectionWithContext(ctx context.Context, request *tcwaf.ModifySpartaProtectionRequest) (response *tcwaf.ModifySpartaProtectionResponse, err error) {
	if request == nil {
		request = tcwaf.NewModifySpartaProtectionRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "waf", tcwaf.APIVersion, "ModifySpartaProtection")

	if c.GetCredential() == nil {
		return nil, errors.New("ModifySpartaProtection require credential")
	}

	request.SetContext(ctx)
	response = tcwaf.NewModifySpartaProtectionResponse()
	err = c.Send(request, response)
	return
}
