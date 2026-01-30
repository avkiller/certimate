package internal

import (
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	alinlb "github.com/alibabacloud-go/nlb-20220430/v4/client"
	"github.com/alibabacloud-go/tea/dara"
)

// This is a partial copy of https://github.com/alibabacloud-go/nlb-20220430/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type NlbClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewNlbClient(config *openapiutil.Config) (*NlbClient, error) {
	client := new(NlbClient)
	err := client.Init(config)
	return client, err
}

func (client *NlbClient) Init(config *openapiutil.Config) (_err error) {
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

func (client *NlbClient) GetListenerAttributeWithContext(ctx context.Context, request *alinlb.GetListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *alinlb.GetListenerAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ListenerId) {
		query["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetListenerAttribute"),
		Version:     dara.String("2022-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alinlb.GetListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *NlbClient) GetLoadBalancerAttributeWithContext(ctx context.Context, request *alinlb.GetLoadBalancerAttributeRequest, runtime *dara.RuntimeOptions) (_result *alinlb.GetLoadBalancerAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoadBalancerAttribute"),
		Version:     dara.String("2022-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alinlb.GetLoadBalancerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *NlbClient) ListListenersWithContext(ctx context.Context, request *alinlb.ListListenersRequest, runtime *dara.RuntimeOptions) (_result *alinlb.ListListenersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.ListenerIds) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerIds) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecSensorEnabled) {
		query["SecSensorEnabled"] = request.SecSensorEnabled
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListeners"),
		Version:     dara.String("2022-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alinlb.ListListenersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *NlbClient) UpdateListenerAttributeWithContext(ctx context.Context, tmpReq *alinlb.UpdateListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *alinlb.UpdateListenerAttributeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}

	request := &alinlb.UpdateListenerAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProxyProtocolV2Config) {
		request.ProxyProtocolV2ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProxyProtocolV2Config, dara.String("ProxyProtocolV2Config"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlpnEnabled) {
		body["AlpnEnabled"] = request.AlpnEnabled
	}

	if !dara.IsNil(request.AlpnPolicy) {
		body["AlpnPolicy"] = request.AlpnPolicy
	}

	if !dara.IsNil(request.CaCertificateIds) {
		body["CaCertificateIds"] = request.CaCertificateIds
	}

	if !dara.IsNil(request.CaEnabled) {
		body["CaEnabled"] = request.CaEnabled
	}

	if !dara.IsNil(request.CertificateIds) {
		body["CertificateIds"] = request.CertificateIds
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Cps) {
		body["Cps"] = request.Cps
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IdleTimeout) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerDescription) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !dara.IsNil(request.ListenerId) {
		body["ListenerId"] = request.ListenerId
	}

	if !dara.IsNil(request.Mss) {
		body["Mss"] = request.Mss
	}

	if !dara.IsNil(request.ProxyProtocolEnabled) {
		body["ProxyProtocolEnabled"] = request.ProxyProtocolEnabled
	}

	if !dara.IsNil(request.ProxyProtocolV2ConfigShrink) {
		body["ProxyProtocolV2Config"] = request.ProxyProtocolV2ConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecSensorEnabled) {
		body["SecSensorEnabled"] = request.SecSensorEnabled
	}

	if !dara.IsNil(request.SecurityPolicyId) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !dara.IsNil(request.ServerGroupId) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateListenerAttribute"),
		Version:     dara.String("2022-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alinlb.UpdateListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
