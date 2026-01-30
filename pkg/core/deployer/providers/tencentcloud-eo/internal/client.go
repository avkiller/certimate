package internal

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcteo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/teo/v20220901/client.go
// to lightweight the vendor packages in the built binary.
type TeoClient struct {
	common.Client
}

func NewTeoClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *TeoClient, err error) {
	client = &TeoClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *TeoClient) DescribeAccelerationDomains(request *tcteo.DescribeAccelerationDomainsRequest) (response *tcteo.DescribeAccelerationDomainsResponse, err error) {
	return c.DescribeAccelerationDomainsWithContext(context.Background(), request)
}

func (c *TeoClient) DescribeAccelerationDomainsWithContext(ctx context.Context, request *tcteo.DescribeAccelerationDomainsRequest) (response *tcteo.DescribeAccelerationDomainsResponse, err error) {
	if request == nil {
		request = tcteo.NewDescribeAccelerationDomainsRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "teo", tcteo.APIVersion, "DescribeAccelerationDomains")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeAccelerationDomains require credential")
	}

	request.SetContext(ctx)
	response = tcteo.NewDescribeAccelerationDomainsResponse()
	err = c.Send(request, response)
	return
}

func (c *TeoClient) ModifyHostsCertificate(request *tcteo.ModifyHostsCertificateRequest) (response *tcteo.ModifyHostsCertificateResponse, err error) {
	return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

func (c *TeoClient) ModifyHostsCertificateWithContext(ctx context.Context, request *tcteo.ModifyHostsCertificateRequest) (response *tcteo.ModifyHostsCertificateResponse, err error) {
	if request == nil {
		request = tcteo.NewModifyHostsCertificateRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "teo", tcteo.APIVersion, "ModifyHostsCertificate")

	if c.GetCredential() == nil {
		return nil, errors.New("ModifyHostsCertificate require credential")
	}

	request.SetContext(ctx)
	response = tcteo.NewModifyHostsCertificateResponse()
	err = c.Send(request, response)
	return
}
