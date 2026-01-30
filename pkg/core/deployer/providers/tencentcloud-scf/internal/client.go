package internal

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcscf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf/v20180416"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/scf/v20180416/client.go
// to lightweight the vendor packages in the built binary.
type ScfClient struct {
	common.Client
}

func NewScfClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *ScfClient, err error) {
	client = &ScfClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *ScfClient) GetCustomDomain(request *tcscf.GetCustomDomainRequest) (response *tcscf.GetCustomDomainResponse, err error) {
	return c.GetCustomDomainWithContext(context.Background(), request)
}

func (c *ScfClient) GetCustomDomainWithContext(ctx context.Context, request *tcscf.GetCustomDomainRequest) (response *tcscf.GetCustomDomainResponse, err error) {
	if request == nil {
		request = tcscf.NewGetCustomDomainRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "scf", tcscf.APIVersion, "GetCustomDomain")

	if c.GetCredential() == nil {
		return nil, errors.New("GetCustomDomain require credential")
	}

	request.SetContext(ctx)
	response = tcscf.NewGetCustomDomainResponse()
	err = c.Send(request, response)
	return
}

func (c *ScfClient) ListCustomDomains(request *tcscf.ListCustomDomainsRequest) (response *tcscf.ListCustomDomainsResponse, err error) {
	return c.ListCustomDomainsWithContext(context.Background(), request)
}

func (c *ScfClient) ListCustomDomainsWithContext(ctx context.Context, request *tcscf.ListCustomDomainsRequest) (response *tcscf.ListCustomDomainsResponse, err error) {
	if request == nil {
		request = tcscf.NewListCustomDomainsRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "scf", tcscf.APIVersion, "ListCustomDomains")

	if c.GetCredential() == nil {
		return nil, errors.New("ListCustomDomains require credential")
	}

	request.SetContext(ctx)

	response = tcscf.NewListCustomDomainsResponse()
	err = c.Send(request, response)
	return
}

func (c *ScfClient) UpdateCustomDomain(request *tcscf.UpdateCustomDomainRequest) (response *tcscf.UpdateCustomDomainResponse, err error) {
	return c.UpdateCustomDomainWithContext(context.Background(), request)
}

func (c *ScfClient) UpdateCustomDomainWithContext(ctx context.Context, request *tcscf.UpdateCustomDomainRequest) (response *tcscf.UpdateCustomDomainResponse, err error) {
	if request == nil {
		request = tcscf.NewUpdateCustomDomainRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "scf", tcscf.APIVersion, "UpdateCustomDomain")

	if c.GetCredential() == nil {
		return nil, errors.New("UpdateCustomDomain require credential")
	}

	request.SetContext(ctx)
	response = tcscf.NewUpdateCustomDomainResponse()
	err = c.Send(request, response)
	return
}
