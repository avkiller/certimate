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

func (c *SslClient) UploadCertificate(request *tcssl.UploadCertificateRequest) (response *tcssl.UploadCertificateResponse, err error) {
	return c.UploadCertificateWithContext(context.Background(), request)
}

func (c *SslClient) UploadCertificateWithContext(ctx context.Context, request *tcssl.UploadCertificateRequest) (response *tcssl.UploadCertificateResponse, err error) {
	if request == nil {
		request = tcssl.NewUploadCertificateRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "ssl", tcssl.APIVersion, "UploadCertificate")

	if c.GetCredential() == nil {
		return nil, errors.New("UploadCertificate require credential")
	}

	request.SetContext(ctx)
	response = tcssl.NewUploadCertificateResponse()
	err = c.Send(request, response)
	return
}
