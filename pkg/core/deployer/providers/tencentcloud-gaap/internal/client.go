package internal

import (
	"context"
	"errors"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcgaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/gaap/v20180529/client.go
// to lightweight the vendor packages in the built binary.
type GaapClient struct {
	common.Client
}

func NewGaapClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *GaapClient, err error) {
	client = &GaapClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *GaapClient) DescribeHTTPSListeners(request *tcgaap.DescribeHTTPSListenersRequest) (response *tcgaap.DescribeHTTPSListenersResponse, err error) {
	return c.DescribeHTTPSListenersWithContext(context.Background(), request)
}

func (c *GaapClient) DescribeHTTPSListenersWithContext(ctx context.Context, request *tcgaap.DescribeHTTPSListenersRequest) (response *tcgaap.DescribeHTTPSListenersResponse, err error) {
	if request == nil {
		request = tcgaap.NewDescribeHTTPSListenersRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeHTTPSListeners require credential")
	}

	request.SetContext(ctx)
	response = tcgaap.NewDescribeHTTPSListenersResponse()
	err = c.Send(request, response)
	return
}

func (c *GaapClient) ModifyHTTPSListenerAttribute(request *tcgaap.ModifyHTTPSListenerAttributeRequest) (response *tcgaap.ModifyHTTPSListenerAttributeResponse, err error) {
	return c.ModifyHTTPSListenerAttributeWithContext(context.Background(), request)
}

func (c *GaapClient) ModifyHTTPSListenerAttributeWithContext(ctx context.Context, request *tcgaap.ModifyHTTPSListenerAttributeRequest) (response *tcgaap.ModifyHTTPSListenerAttributeResponse, err error) {
	if request == nil {
		request = tcgaap.NewModifyHTTPSListenerAttributeRequest()
	}

	if c.GetCredential() == nil {
		return nil, errors.New("ModifyHTTPSListenerAttribute require credential")
	}

	request.SetContext(ctx)
	response = tcgaap.NewModifyHTTPSListenerAttributeResponse()
	err = c.Send(request, response)
	return
}
