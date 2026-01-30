package internal

import (
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	aliddoscoo "github.com/alibabacloud-go/ddoscoo-20200101/v4/client"
	"github.com/alibabacloud-go/tea/dara"
)

// This is a partial copy of https://github.com/alibabacloud-go/ddoscoo-20200101/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type DdoscooClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewDdoscooClient(config *openapiutil.Config) (*DdoscooClient, error) {
	client := new(DdoscooClient)
	err := client.Init(config)
	return client, err
}

func (client *DdoscooClient) Init(config *openapiutil.Config) (_err error) {
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

func (client *DdoscooClient) AssociateWebCertWithContext(ctx context.Context, request *aliddoscoo.AssociateWebCertRequest, runtime *dara.RuntimeOptions) (_result *aliddoscoo.AssociateWebCertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}

	if !dara.IsNil(request.Cert) {
		body["Cert"] = request.Cert
	}

	if !dara.IsNil(request.CertId) {
		body["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.CertName) {
		body["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		body["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.Domain) {
		body["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Key) {
		body["Key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateWebCert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &aliddoscoo.AssociateWebCertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *DdoscooClient) DescribeDomainsWithContext(ctx context.Context, request *aliddoscoo.DescribeDomainsRequest, runtime *dara.RuntimeOptions) (_result *aliddoscoo.DescribeDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomains"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &aliddoscoo.DescribeDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
