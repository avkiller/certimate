package internal

import (
	"github.com/volcengine/volcengine-go-sdk/service/certificateservice"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginequery"
)

// This is a partial copy of https://github.com/volcengine/volcengine-go-sdk/blob/master/service/certificateservice/service_certificateservice.go
// to lightweight the vendor packages in the built binary.
type CertificateserviceClient struct {
	*client.Client
}

func NewCertificateserviceClient(p client.ConfigProvider, cfgs ...*volcengine.Config) *CertificateserviceClient {
	c := p.ClientConfig(certificateservice.EndpointsID, cfgs...)
	return newCertificateserviceClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

func newCertificateserviceClient(cfg volcengine.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CertificateserviceClient {
	svc := &CertificateserviceClient{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   certificateservice.ServiceName,
				ServiceID:     certificateservice.ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2024-10-01",
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

func (c *CertificateserviceClient) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

func (c *CertificateserviceClient) ImportCertificate(input *certificateservice.ImportCertificateInput) (*certificateservice.ImportCertificateOutput, error) {
	req, out := c.ImportCertificateRequest(input)
	return out, req.Send()
}

func (c *CertificateserviceClient) ImportCertificateRequest(input *certificateservice.ImportCertificateInput) (req *request.Request, output *certificateservice.ImportCertificateOutput) {
	op := &request.Operation{
		Name:       "ImportCertificate",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &certificateservice.ImportCertificateInput{}
	}

	output = &certificateservice.ImportCertificateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}
