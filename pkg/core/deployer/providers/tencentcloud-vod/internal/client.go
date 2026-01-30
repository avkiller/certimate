package internal

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcvod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/vod/v20180717/client.go
// to lightweight the vendor packages in the built binary.
type VodClient struct {
	common.Client
}

func NewVodClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *VodClient, err error) {
	client = &VodClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *VodClient) DescribeVodDomains(request *tcvod.DescribeVodDomainsRequest) (response *tcvod.DescribeVodDomainsResponse, err error) {
	return c.DescribeVodDomainsWithContext(context.Background(), request)
}

func (c *VodClient) DescribeVodDomainsWithContext(ctx context.Context, request *tcvod.DescribeVodDomainsRequest) (response *tcvod.DescribeVodDomainsResponse, err error) {
	if request == nil {
		request = tcvod.NewDescribeVodDomainsRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "vod", tcvod.APIVersion, "DescribeVodDomains")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeVodDomains require credential")
	}

	request.SetContext(ctx)

	response = tcvod.NewDescribeVodDomainsResponse()
	err = c.Send(request, response)
	return
}

func (c *VodClient) SetVodDomainCertificate(request *tcvod.SetVodDomainCertificateRequest) (response *tcvod.SetVodDomainCertificateResponse, err error) {
	return c.SetVodDomainCertificateWithContext(context.Background(), request)
}

func (c *VodClient) SetVodDomainCertificateWithContext(ctx context.Context, request *tcvod.SetVodDomainCertificateRequest) (response *tcvod.SetVodDomainCertificateResponse, err error) {
	if request == nil {
		request = tcvod.NewSetVodDomainCertificateRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "vod", tcvod.APIVersion, "SetVodDomainCertificate")

	if c.GetCredential() == nil {
		return nil, errors.New("SetVodDomainCertificate require credential")
	}

	request.SetContext(ctx)
	response = tcvod.NewSetVodDomainCertificateResponse()
	err = c.Send(request, response)
	return
}
