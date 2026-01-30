package internal

import (
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutilv2 "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	alifc "github.com/alibabacloud-go/fc-20230330/v4/client"
	alifcopen "github.com/alibabacloud-go/fc-open-20210406/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
)

// This is a partial copy of https://github.com/alibabacloud-go/fc-20230330/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type FcClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewFcClient(config *openapiutilv2.Config) (*FcClient, error) {
	client := new(FcClient)
	err := client.Init(config)
	return client, err
}

func (client *FcClient) Init(config *openapiutilv2.Config) (_err error) {
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

func (client *FcClient) GetCustomDomainWithContext(ctx context.Context, domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *alifc.GetCustomDomainResponse, _err error) {
	req := &openapiutilv2.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutilv2.Params{
		Action:      dara.String("GetCustomDomain"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &alifc.GetCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *FcClient) ListCustomDomainsWithContext(ctx context.Context, request *alifc.ListCustomDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *alifc.ListCustomDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Prefix) {
		query["prefix"] = request.Prefix
	}

	req := &openapiutilv2.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutilv2.Params{
		Action:      dara.String("ListCustomDomains"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &alifc.ListCustomDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *FcClient) UpdateCustomDomainWithContext(ctx context.Context, domainName *string, request *alifc.UpdateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *alifc.UpdateCustomDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}

	req := &openapiutilv2.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutilv2.Params{
		Action:      dara.String("UpdateCustomDomain"),
		Version:     dara.String("2023-03-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2023-03-30/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &alifc.UpdateCustomDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// This is a partial copy of https://github.com/alibabacloud-go/fc-open-20210406/blob/master/client/client.go
// to lightweight the vendor packages in the built binary.
type FcopenClient struct {
	openapi.Client
}

func NewFcopenClient(config *openapi.Config) (*FcopenClient, error) {
	client := new(FcopenClient)
	err := client.Init(config)
	return client, err
}

func (client *FcopenClient) Init(config *openapi.Config) (_err error) {
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

func (client *FcopenClient) GetCustomDomain(domainName *string) (_result *alifcopen.GetCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &alifcopen.GetCustomDomainHeaders{}
	_result = &alifcopen.GetCustomDomainResponse{}
	_body, _err := client.GetCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *FcopenClient) GetCustomDomainWithOptions(domainName *string, headers *alifcopen.GetCustomDomainHeaders, runtime *util.RuntimeOptions) (_result *alifcopen.GetCustomDomainResponse, _err error) {
	realHeaders := make(map[string]*string)

	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomDomain"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &alifcopen.GetCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *FcopenClient) ListCustomDomains(request *alifcopen.ListCustomDomainsRequest) (_result *alifcopen.ListCustomDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &alifcopen.ListCustomDomainsHeaders{}
	_result = &alifcopen.ListCustomDomainsResponse{}
	_body, _err := client.ListCustomDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *FcopenClient) ListCustomDomainsWithOptions(request *alifcopen.ListCustomDomainsRequest, headers *alifcopen.ListCustomDomainsHeaders, runtime *util.RuntimeOptions) (_result *alifcopen.ListCustomDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartKey)) {
		query["startKey"] = request.StartKey
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomDomains"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &alifcopen.ListCustomDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *FcopenClient) UpdateCustomDomain(domainName *string, request *alifcopen.UpdateCustomDomainRequest) (_result *alifcopen.UpdateCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &alifcopen.UpdateCustomDomainHeaders{}
	_result = &alifcopen.UpdateCustomDomainResponse{}
	_body, _err := client.UpdateCustomDomainWithOptions(domainName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *FcopenClient) UpdateCustomDomainWithOptions(domainName *string, request *alifcopen.UpdateCustomDomainRequest, headers *alifcopen.UpdateCustomDomainHeaders, runtime *util.RuntimeOptions) (_result *alifcopen.UpdateCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.CertConfig)) {
		body["certConfig"] = request.CertConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RouteConfig)) {
		body["routeConfig"] = request.RouteConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TlsConfig)) {
		body["tlsConfig"] = request.TlsConfig
	}

	if !tea.BoolValue(util.IsUnset(request.WafConfig)) {
		body["wafConfig"] = request.WafConfig
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAccountId)) {
		realHeaders["X-Fc-Account-Id"] = util.ToJSONString(headers.XFcAccountId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcDate)) {
		realHeaders["X-Fc-Date"] = util.ToJSONString(headers.XFcDate)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcTraceId)) {
		realHeaders["X-Fc-Trace-Id"] = util.ToJSONString(headers.XFcTraceId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomDomain"),
		Version:     tea.String("2021-04-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2021-04-06/custom-domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &alifcopen.UpdateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
