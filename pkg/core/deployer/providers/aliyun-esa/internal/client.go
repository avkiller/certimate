package internal

import (
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	aliesa "github.com/alibabacloud-go/esa-20240910/v2/client"
	"github.com/alibabacloud-go/tea/dara"
)

// This is a partial copy of https://github.com/alibabacloud-go/esa-20240910/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type EsaClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewEsaClient(config *openapiutil.Config) (*EsaClient, error) {
	client := new(EsaClient)
	err := client.Init(config)
	return client, err
}

func (client *EsaClient) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *EsaClient) SetCertificateWithContext(ctx context.Context, request *aliesa.SetCertificateRequest, runtime *dara.RuntimeOptions) (_result *aliesa.SetCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CasId) {
		body["CasId"] = request.CasId
	}

	if !dara.IsNil(request.Certificate) {
		body["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PrivateKey) {
		body["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCertificate"),
		Version:     dara.String("2024-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &aliesa.SetCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
