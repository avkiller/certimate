package internal

import (
	"context"
	"errors"

	tccdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/cdn/v20180606/client.go
// to lightweight the vendor packages in the built binary.
type CdnClient struct {
	common.Client
}

func NewCdnClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *CdnClient, err error) {
	client = &CdnClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *CdnClient) DescribeCertDomains(request *tccdn.DescribeCertDomainsRequest) (response *tccdn.DescribeCertDomainsResponse, err error) {
	return c.DescribeCertDomainsWithContext(context.Background(), request)
}

func (c *CdnClient) DescribeCertDomainsWithContext(ctx context.Context, request *tccdn.DescribeCertDomainsRequest) (response *tccdn.DescribeCertDomainsResponse, err error) {
	if request == nil {
		request = tccdn.NewDescribeCertDomainsRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "cdn", tccdn.APIVersion, "DescribeCertDomains")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeCertDomains require credential")
	}

	request.SetContext(ctx)
	response = tccdn.NewDescribeCertDomainsResponse()
	err = c.Send(request, response)
	return
}

func (c *CdnClient) DescribeDomains(request *tccdn.DescribeDomainsRequest) (response *tccdn.DescribeDomainsResponse, err error) {
	return c.DescribeDomainsWithContext(context.Background(), request)
}

func (c *CdnClient) DescribeDomainsWithContext(ctx context.Context, request *tccdn.DescribeDomainsRequest) (response *tccdn.DescribeDomainsResponse, err error) {
	if request == nil {
		request = tccdn.NewDescribeDomainsRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "cdn", tccdn.APIVersion, "DescribeDomains")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeDomains require credential")
	}

	request.SetContext(ctx)
	response = tccdn.NewDescribeDomainsResponse()
	err = c.Send(request, response)
	return
}

func (c *CdnClient) DescribeDomainsConfig(request *tccdn.DescribeDomainsConfigRequest) (response *tccdn.DescribeDomainsConfigResponse, err error) {
	return c.DescribeDomainsConfigWithContext(context.Background(), request)
}

func (c *CdnClient) DescribeDomainsConfigWithContext(ctx context.Context, request *tccdn.DescribeDomainsConfigRequest) (response *tccdn.DescribeDomainsConfigResponse, err error) {
	if request == nil {
		request = tccdn.NewDescribeDomainsConfigRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "cdn", tccdn.APIVersion, "DescribeDomainsConfig")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeDomainsConfig require credential")
	}

	request.SetContext(ctx)
	response = tccdn.NewDescribeDomainsConfigResponse()
	err = c.Send(request, response)
	return
}

func (c *CdnClient) UpdateDomainConfig(request *tccdn.UpdateDomainConfigRequest) (response *tccdn.UpdateDomainConfigResponse, err error) {
	return c.UpdateDomainConfigWithContext(context.Background(), request)
}

func (c *CdnClient) UpdateDomainConfigWithContext(ctx context.Context, request *tccdn.UpdateDomainConfigRequest) (response *tccdn.UpdateDomainConfigResponse, err error) {
	if request == nil {
		request = tccdn.NewUpdateDomainConfigRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "cdn", tccdn.APIVersion, "UpdateDomainConfig")

	if c.GetCredential() == nil {
		return nil, errors.New("UpdateDomainConfig require credential")
	}

	request.SetContext(ctx)
	response = tccdn.NewUpdateDomainConfigResponse()
	err = c.Send(request, response)
	return
}
