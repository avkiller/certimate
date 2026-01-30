package internal

import (
	"github.com/volcengine/volcengine-go-sdk/service/cdn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

// This is a partial copy of https://github.com/volcengine/volcengine-go-sdk/blob/master/service/cdn/service_cdn.go
// to lightweight the vendor packages in the built binary.
type CdnClient struct {
	*client.Client
}

func NewCdnClient(p client.ConfigProvider, cfgs ...*volcengine.Config) *CdnClient {
	c := p.ClientConfig(cdn.EndpointsID, cfgs...)
	return newCdnClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

func newCdnClient(cfg volcengine.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CdnClient {
	svc := &CdnClient{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   cdn.ServiceName,
				ServiceID:     cdn.ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2021-03-01",
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

func (c *CdnClient) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

func (c *CdnClient) BatchDeployCert(input *cdn.BatchDeployCertInput) (*cdn.BatchDeployCertOutput, error) {
	req, out := c.BatchDeployCertRequest(input)
	return out, req.Send()
}

func (c *CdnClient) BatchDeployCertRequest(input *cdn.BatchDeployCertInput) (req *request.Request, output *cdn.BatchDeployCertOutput) {
	op := &request.Operation{
		Name:       "BatchDeployCert",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &cdn.BatchDeployCertInput{}
	}

	output = &cdn.BatchDeployCertOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

func (c *CdnClient) DescribeCertConfig(input *cdn.DescribeCertConfigInput) (*cdn.DescribeCertConfigOutput, error) {
	req, out := c.DescribeCertConfigRequest(input)
	return out, req.Send()
}

func (c *CdnClient) DescribeCertConfigRequest(input *cdn.DescribeCertConfigInput) (req *request.Request, output *cdn.DescribeCertConfigOutput) {
	op := &request.Operation{
		Name:       "DescribeCertConfig",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &cdn.DescribeCertConfigInput{}
	}

	output = &cdn.DescribeCertConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

func (c *CdnClient) ListCdnDomains(input *cdn.ListCdnDomainsInput) (*cdn.ListCdnDomainsOutput, error) {
	req, out := c.ListCdnDomainsRequest(input)
	return out, req.Send()
}

func (c *CdnClient) ListCdnDomainsRequest(input *cdn.ListCdnDomainsInput) (req *request.Request, output *cdn.ListCdnDomainsOutput) {
	op := &request.Operation{
		Name:       "ListCdnDomains",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &cdn.ListCdnDomainsInput{}
	}

	output = &cdn.ListCdnDomainsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}
