package internal

import (
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

// This is a partial copy of https://github.com/volcengine/volcengine-go-sdk/blob/master/service/clb/service_clb.go
// to lightweight the vendor packages in the built binary.
type ClbClient struct {
	*client.Client
}

func NewClbClient(p client.ConfigProvider, cfgs ...*volcengine.Config) *ClbClient {
	c := p.ClientConfig(clb.EndpointsID, cfgs...)
	return newClbClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

func newClbClient(cfg volcengine.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ClbClient {
	svc := &ClbClient{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   clb.ServiceName,
				ServiceID:     clb.ServiceID,
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

func (c *ClbClient) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

func (c *ClbClient) DescribeListeners(input *clb.DescribeListenersInput) (*clb.DescribeListenersOutput, error) {
	req, out := c.DescribeListenersRequest(input)
	return out, req.Send()
}

func (c *ClbClient) DescribeListenersRequest(input *clb.DescribeListenersInput) (req *request.Request, output *clb.DescribeListenersOutput) {
	op := &request.Operation{
		Name:       "DescribeListeners",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &clb.DescribeListenersInput{}
	}

	output = &clb.DescribeListenersOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *ClbClient) DescribeLoadBalancerAttributes(input *clb.DescribeLoadBalancerAttributesInput) (*clb.DescribeLoadBalancerAttributesOutput, error) {
	req, out := c.DescribeLoadBalancerAttributesRequest(input)
	return out, req.Send()
}

func (c *ClbClient) DescribeLoadBalancerAttributesRequest(input *clb.DescribeLoadBalancerAttributesInput) (req *request.Request, output *clb.DescribeLoadBalancerAttributesOutput) {
	op := &request.Operation{
		Name:       "DescribeLoadBalancerAttributes",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &clb.DescribeLoadBalancerAttributesInput{}
	}

	output = &clb.DescribeLoadBalancerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

func (c *ClbClient) ModifyListenerAttributes(input *clb.ModifyListenerAttributesInput) (*clb.ModifyListenerAttributesOutput, error) {
	req, out := c.ModifyListenerAttributesRequest(input)
	return out, req.Send()
}

func (c *ClbClient) ModifyListenerAttributesRequest(input *clb.ModifyListenerAttributesInput) (req *request.Request, output *clb.ModifyListenerAttributesOutput) {
	op := &request.Operation{
		Name:       "ModifyListenerAttributes",
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &clb.ModifyListenerAttributesInput{}
	}

	output = &clb.ModifyListenerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}
