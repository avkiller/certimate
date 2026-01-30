package internal

import (
	"github.com/volcengine/volcengine-go-sdk/service/alb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

// This is a partial copy of https://github.com/volcengine/volcengine-go-sdk/blob/master/service/alb/service_alb.go
// to lightweight the vendor packages in the built binary.
type AlbClient struct {
	*client.Client
}

func NewAlbClient(p client.ConfigProvider, cfgs ...*volcengine.Config) *AlbClient {
	c := p.ClientConfig(alb.EndpointsID, cfgs...)
	return newAlbClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

func newAlbClient(cfg volcengine.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *AlbClient {
	svc := &AlbClient{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   alb.ServiceName,
				ServiceID:     alb.ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2020-04-01",
			},
			handlers,
		),
	}

	svc.Handlers.Build.PushBackNamed(corehandlers.SDKVersionUserAgentHandler)
	svc.Handlers.Build.PushBackNamed(corehandlers.AddHostExecEnvUserAgentHandler)
	svc.Handlers.Sign.PushBackNamed(volc.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(volcenginequery.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(volcenginequery.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(volcenginequery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(volcenginequery.UnmarshalErrorHandler)

	return svc
}

func (c *AlbClient) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

func (c *AlbClient) DescribeListenerAttributes(input *alb.DescribeListenerAttributesInput) (*alb.DescribeListenerAttributesOutput, error) {
	req, out := c.DescribeListenerAttributesRequest(input)
	return out, req.Send()
}

func (c *AlbClient) DescribeListenerAttributesRequest(input *alb.DescribeListenerAttributesInput) (req *request.Request, output *alb.DescribeListenerAttributesOutput) {
	op := &request.Operation{
		Name:       "DescribeListenerAttributes",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &alb.DescribeListenerAttributesInput{}
	}

	output = &alb.DescribeListenerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *AlbClient) DescribeListeners(input *alb.DescribeListenersInput) (*alb.DescribeListenersOutput, error) {
	req, out := c.DescribeListenersRequest(input)
	return out, req.Send()
}

func (c *AlbClient) DescribeListenersRequest(input *alb.DescribeListenersInput) (req *request.Request, output *alb.DescribeListenersOutput) {
	op := &request.Operation{
		Name:       "DescribeListeners",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &alb.DescribeListenersInput{}
	}

	output = &alb.DescribeListenersOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *AlbClient) DescribeLoadBalancerAttributes(input *alb.DescribeLoadBalancerAttributesInput) (*alb.DescribeLoadBalancerAttributesOutput, error) {
	req, out := c.DescribeLoadBalancerAttributesRequest(input)
	return out, req.Send()
}

func (c *AlbClient) DescribeLoadBalancerAttributesRequest(input *alb.DescribeLoadBalancerAttributesInput) (req *request.Request, output *alb.DescribeLoadBalancerAttributesOutput) {
	op := &request.Operation{
		Name:       "DescribeLoadBalancerAttributes",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &alb.DescribeLoadBalancerAttributesInput{}
	}

	output = &alb.DescribeLoadBalancerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *AlbClient) ModifyListenerAttributes(input *alb.ModifyListenerAttributesInput) (*alb.ModifyListenerAttributesOutput, error) {
	req, out := c.ModifyListenerAttributesRequest(input)
	return out, req.Send()
}

func (c *AlbClient) ModifyListenerAttributesRequest(input *alb.ModifyListenerAttributesInput) (req *request.Request, output *alb.ModifyListenerAttributesOutput) {
	op := &request.Operation{
		Name:       "ModifyListenerAttributes",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &alb.ModifyListenerAttributesInput{}
	}

	output = &alb.ModifyListenerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}
