package internal

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/ssl/v20191205/client.go
// to lightweight the vendor packages in the built binary.
type SslClient struct {
	common.Client
}

func NewSslClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *SslClient, err error) {
	client = &SslClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *SslClient) DescribeHostCosInstanceList(request *tcssl.DescribeHostCosInstanceListRequest) (response *tcssl.DescribeHostCosInstanceListResponse, err error) {
	return c.DescribeHostCosInstanceListWithContext(context.Background(), request)
}

func (c *SslClient) DescribeHostCosInstanceListWithContext(ctx context.Context, request *tcssl.DescribeHostCosInstanceListRequest) (response *tcssl.DescribeHostCosInstanceListResponse, err error) {
	if request == nil {
		request = tcssl.NewDescribeHostCosInstanceListRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "DescribeHostCosInstanceList")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeHostCosInstanceList require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewDescribeHostCosInstanceListResponse()
	err = c.Send(request, response)
	return
}

func (c *SslClient) DescribeHostDeployRecordDetail(request *tcssl.DescribeHostDeployRecordDetailRequest) (response *tcssl.DescribeHostDeployRecordDetailResponse, err error) {
	return c.DescribeHostDeployRecordDetailWithContext(context.Background(), request)
}

func (c *SslClient) DescribeHostDeployRecordDetailWithContext(ctx context.Context, request *tcssl.DescribeHostDeployRecordDetailRequest) (response *tcssl.DescribeHostDeployRecordDetailResponse, err error) {
	if request == nil {
		request = tcssl.NewDescribeHostDeployRecordDetailRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "DescribeHostDeployRecordDetail")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeHostDeployRecordDetail require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewDescribeHostDeployRecordDetailResponse()
	err = c.Send(request, response)
	return
}

func (c *SslClient) DeployCertificateInstance(request *tcssl.DeployCertificateInstanceRequest) (response *tcssl.DeployCertificateInstanceResponse, err error) {
	return c.DeployCertificateInstanceWithContext(context.Background(), request)
}

func (c *SslClient) DeployCertificateInstanceWithContext(ctx context.Context, request *tcssl.DeployCertificateInstanceRequest) (response *tcssl.DeployCertificateInstanceResponse, err error) {
	if request == nil {
		request = tcssl.NewDeployCertificateInstanceRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "DeployCertificateInstance")

	if c.GetCredential() == nil {
		return nil, errors.New("DeployCertificateInstance require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewDeployCertificateInstanceResponse()
	err = c.Send(request, response)
	return
}
