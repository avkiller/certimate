package internal

import (
	"context"

	aliapig "github.com/alibabacloud-go/apig-20240327/v5/client"
	alicloudapi "github.com/alibabacloud-go/cloudapi-20160714/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// This is a partial copy of https://github.com/alibabacloud-go/apig-20240327/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type ApigClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewApigClient(config *openapiutil.Config) (*ApigClient, error) {
	client := new(ApigClient)
	err := client.Init(config)
	return client, err
}

func (client *ApigClient) GetDomainWithContext(ctx context.Context, domainId *string, request *aliapig.GetDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *aliapig.GetDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.WithStatistics) {
		query["withStatistics"] = request.WithStatistics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &aliapig.GetDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *ApigClient) ListDomainsWithContext(ctx context.Context, request *aliapig.ListDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *aliapig.ListDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &aliapig.ListDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *ApigClient) UpdateDomainWithContext(ctx context.Context, domainId *string, request *aliapig.UpdateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *aliapig.UpdateDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}

	if !dara.IsNil(request.CaCertIdentifier) {
		body["caCertIdentifier"] = request.CaCertIdentifier
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["certIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.ClientCACert) {
		body["clientCACert"] = request.ClientCACert
	}

	if !dara.IsNil(request.ForceHttps) {
		body["forceHttps"] = request.ForceHttps
	}

	if !dara.IsNil(request.Http2Option) {
		body["http2Option"] = request.Http2Option
	}

	if !dara.IsNil(request.MTLSEnabled) {
		body["mTLSEnabled"] = request.MTLSEnabled
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TlsCipherSuitesConfig) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
	}

	if !dara.IsNil(request.TlsMax) {
		body["tlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &aliapig.UpdateDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// This is a partial copy of https://github.com/alibabacloud-go/cloudapi-20160714/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type CloudapiClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewCloudapiClient(config *openapiutil.Config) (*CloudapiClient, error) {
	client := new(CloudapiClient)
	err := client.Init(config)
	return client, err
}

func (client *CloudapiClient) DescribeApiGroupWithContext(ctx context.Context, request *alicloudapi.DescribeApiGroupRequest, runtime *dara.RuntimeOptions) (_result *alicloudapi.DescribeApiGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alicloudapi.DescribeApiGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *CloudapiClient) SetDomainCertificateWithContext(ctx context.Context, request *alicloudapi.SetDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *alicloudapi.SetDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.CaCertificateBody) {
		query["CaCertificateBody"] = request.CaCertificateBody
	}

	if !dara.IsNil(request.CertificateBody) {
		query["CertificateBody"] = request.CertificateBody
	}

	if !dara.IsNil(request.CertificateName) {
		query["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.CertificatePrivateKey) {
		query["CertificatePrivateKey"] = request.CertificatePrivateKey
	}

	if !dara.IsNil(request.ClientCertSDnPassThrough) {
		query["ClientCertSDnPassThrough"] = request.ClientCertSDnPassThrough
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SslOcspCacheEnable) {
		query["SslOcspCacheEnable"] = request.SslOcspCacheEnable
	}

	if !dara.IsNil(request.SslOcspEnable) {
		query["SslOcspEnable"] = request.SslOcspEnable
	}

	if !dara.IsNil(request.SslVerifyDepth) {
		query["SslVerifyDepth"] = request.SslVerifyDepth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainCertificate"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alicloudapi.SetDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
