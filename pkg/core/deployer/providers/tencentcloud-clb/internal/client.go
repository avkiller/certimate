package internal

import (
	"context"
	"errors"

	tcclb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

// This is a partial copy of https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/clb/v20180317/client.go
// to lightweight the vendor packages in the built binary.
type ClbClient struct {
	common.Client
}

func NewClbClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *ClbClient, err error) {
	client = &ClbClient{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func (c *ClbClient) DescribeListeners(request *tcclb.DescribeListenersRequest) (response *tcclb.DescribeListenersResponse, err error) {
	return c.DescribeListenersWithContext(context.Background(), request)
}

func (c *ClbClient) DescribeListenersWithContext(ctx context.Context, request *tcclb.DescribeListenersRequest) (response *tcclb.DescribeListenersResponse, err error) {
	if request == nil {
		request = tcclb.NewDescribeListenersRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "clb", tcclb.APIVersion, "DescribeListeners")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeListeners require credential")
	}

	request.SetContext(ctx)
	response = tcclb.NewDescribeListenersResponse()
	err = c.Send(request, response)
	return
}

func (c *ClbClient) DescribeTaskStatus(request *tcclb.DescribeTaskStatusRequest) (response *tcclb.DescribeTaskStatusResponse, err error) {
	return c.DescribeTaskStatusWithContext(context.Background(), request)
}

func (c *ClbClient) DescribeTaskStatusWithContext(ctx context.Context, request *tcclb.DescribeTaskStatusRequest) (response *tcclb.DescribeTaskStatusResponse, err error) {
	if request == nil {
		request = tcclb.NewDescribeTaskStatusRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "clb", tcclb.APIVersion, "DescribeTaskStatus")

	if c.GetCredential() == nil {
		return nil, errors.New("DescribeTaskStatus require credential")
	}

	request.SetContext(ctx)
	response = tcclb.NewDescribeTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func (c *ClbClient) ModifyDomainAttributes(request *tcclb.ModifyDomainAttributesRequest) (response *tcclb.ModifyDomainAttributesResponse, err error) {
	return c.ModifyDomainAttributesWithContext(context.Background(), request)
}

func (c *ClbClient) ModifyDomainAttributesWithContext(ctx context.Context, request *tcclb.ModifyDomainAttributesRequest) (response *tcclb.ModifyDomainAttributesResponse, err error) {
	if request == nil {
		request = tcclb.NewModifyDomainAttributesRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "clb", tcclb.APIVersion, "ModifyDomainAttributes")

	if c.GetCredential() == nil {
		return nil, errors.New("ModifyDomainAttributes require credential")
	}

	request.SetContext(ctx)
	response = tcclb.NewModifyDomainAttributesResponse()
	err = c.Send(request, response)
	return
}

func (c *ClbClient) ModifyListener(request *tcclb.ModifyListenerRequest) (response *tcclb.ModifyListenerResponse, err error) {
	return c.ModifyListenerWithContext(context.Background(), request)
}

func (c *ClbClient) ModifyListenerWithContext(ctx context.Context, request *tcclb.ModifyListenerRequest) (response *tcclb.ModifyListenerResponse, err error) {
	if request == nil {
		request = tcclb.NewModifyListenerRequest()
	}
	c.InitBaseRequest(&request.BaseRequest, "clb", tcclb.APIVersion, "ModifyListener")

	if c.GetCredential() == nil {
		return nil, errors.New("ModifyListener require credential")
	}

	request.SetContext(ctx)
	response = tcclb.NewModifyListenerResponse()
	err = c.Send(request, response)
	return
}
