package internal

import (
	"github.com/volcengine/volcengine-go-sdk/service/dcdn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

// This is a partial copy of https://github.com/volcengine/volcengine-go-sdk/blob/master/service/dcdn/service_dcdn.go
// to lightweight the vendor packages in the built binary.
type DcdnClient struct {
	*client.Client
}

func NewDcdnClient(p client.ConfigProvider, cfgs ...*volcengine.Config) *DcdnClient {
	c := p.ClientConfig(dcdn.EndpointsID, cfgs...)
	return newDcdnClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

func newDcdnClient(cfg volcengine.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *DcdnClient {
	svc := &DcdnClient{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   dcdn.ServiceName,
				ServiceID:     dcdn.ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2021-04-01",
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

func (c *DcdnClient) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

func (c *DcdnClient) CreateCertBind(input *dcdn.CreateCertBindInput) (*dcdn.CreateCertBindOutput, error) {
	req, out := c.CreateCertBindRequest(input)
	return out, req.Send()
}

func (c *DcdnClient) CreateCertBindRequest(input *dcdn.CreateCertBindInput) (req *request.Request, output *dcdn.CreateCertBindOutput) {
	op := &request.Operation{
		Name:       "CreateCertBind",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &dcdn.CreateCertBindInput{}
	}

	output = &dcdn.CreateCertBindOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

func (c *DcdnClient) ListDomainConfig(input *dcdn.ListDomainConfigInput) (*dcdn.ListDomainConfigOutput, error) {
	req, out := c.ListDomainConfigRequest(input)
	return out, req.Send()
}

func (c *DcdnClient) ListDomainConfigRequest(input *dcdn.ListDomainConfigInput) (req *request.Request, output *dcdn.ListDomainConfigOutput) {
	op := &request.Operation{
		Name:       "ListDomainConfig",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &dcdn.ListDomainConfigInput{}
	}

	output = &dcdn.ListDomainConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}
