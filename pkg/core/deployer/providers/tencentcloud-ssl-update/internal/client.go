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

func (c *SslClient) DescribeHostUpdateRecordDetail(request *tcssl.DescribeHostUpdateRecordDetailRequest) (response *tcssl.DescribeHostUpdateRecordDetailResponse, err error) {
	return c.DescribeHostUpdateRecordDetailWithContext(context.Background(), request)
}

func (c *SslClient) DescribeHostUpdateRecordDetailWithContext(ctx context.Context, request *tcssl.DescribeHostUpdateRecordDetailRequest) (response *tcssl.DescribeHostUpdateRecordDetailResponse, err error) {
	if request == nil {
		request = tcssl.NewDescribeHostUpdateRecordDetailRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "DescribeHostUpdateRecordDetail")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeHostUpdateRecordDetail require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewDescribeHostUpdateRecordDetailResponse()
	err = c.Send(request, response)
	return
}

func (c *SslClient) DescribeHostUploadUpdateRecordDetail(request *tcssl.DescribeHostUploadUpdateRecordDetailRequest) (response *tcssl.DescribeHostUploadUpdateRecordDetailResponse, err error) {
	return c.DescribeHostUploadUpdateRecordDetailWithContext(context.Background(), request)
}

func (c *SslClient) DescribeHostUploadUpdateRecordDetailWithContext(ctx context.Context, request *tcssl.DescribeHostUploadUpdateRecordDetailRequest) (response *tcssl.DescribeHostUploadUpdateRecordDetailResponse, err error) {
	if request == nil {
		request = tcssl.NewDescribeHostUploadUpdateRecordDetailRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "DescribeHostUploadUpdateRecordDetail")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeHostUploadUpdateRecordDetail require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewDescribeHostUploadUpdateRecordDetailResponse()
	err = c.Send(request, response)
	return
}

func (c *SslClient) UpdateCertificateInstance(request *tcssl.UpdateCertificateInstanceRequest) (response *tcssl.UpdateCertificateInstanceResponse, err error) {
	return c.UpdateCertificateInstanceWithContext(context.Background(), request)
}

func (c *SslClient) UpdateCertificateInstanceWithContext(ctx context.Context, request *tcssl.UpdateCertificateInstanceRequest) (response *tcssl.UpdateCertificateInstanceResponse, err error) {
	if request == nil {
		request = tcssl.NewUpdateCertificateInstanceRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "UpdateCertificateInstance")

	if c.GetCredential() == nil {
		return nil, errors.New("UpdateCertificateInstance require credential")
	}

	request.SetContext(ctx)

	response = tcssl.NewUpdateCertificateInstanceResponse()
	err = c.Send(request, response)
	return
}

func (c *SslClient) UploadUpdateCertificateInstance(request *tcssl.UploadUpdateCertificateInstanceRequest) (response *tcssl.UploadUpdateCertificateInstanceResponse, err error) {
	return c.UploadUpdateCertificateInstanceWithContext(context.Background(), request)
}

func (c *SslClient) UploadUpdateCertificateInstanceWithContext(ctx context.Context, request *tcssl.UploadUpdateCertificateInstanceRequest) (response *tcssl.UploadUpdateCertificateInstanceResponse, err error) {
	if request == nil {
		request = tcssl.NewUploadUpdateCertificateInstanceRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "UploadUpdateCertificateInstance")

	if c.GetCredential() == nil {
		return nil, errors.New("UploadUpdateCertificateInstance require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewUploadUpdateCertificateInstanceResponse()
	err = c.Send(request, response)
	return
}
