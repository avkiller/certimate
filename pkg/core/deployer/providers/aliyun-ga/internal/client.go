package internal

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	aliga "github.com/alibabacloud-go/ga-20191120/v3/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// This is a partial copy of https://github.com/alibabacloud-go/ga-20191120/blob/master/client/client.go
// to lightweight the vendor packages in the built binary.
type GaClient struct {
	openapi.Client
}

func NewGaClient(config *openapi.Config) (*GaClient, error) {
	client := new(GaClient)
	err := client.Init(config)
	return client, err
}

func (client *GaClient) Init(config *openapi.Config) (_err error) {
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

func (client *GaClient) AssociateAdditionalCertificatesWithListenerWithOptions(request *aliga.AssociateAdditionalCertificatesWithListenerRequest, runtime *util.RuntimeOptions) (_result *aliga.AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateAdditionalCertificatesWithListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &aliga.AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *GaClient) AssociateAdditionalCertificatesWithListener(request *aliga.AssociateAdditionalCertificatesWithListenerRequest) (_result *aliga.AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &aliga.AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.AssociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *GaClient) ListListenersWithOptions(request *aliga.ListListenersRequest, runtime *util.RuntimeOptions) (_result *aliga.ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListeners"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &aliga.ListListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *GaClient) ListListeners(request *aliga.ListListenersRequest) (_result *aliga.ListListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &aliga.ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *GaClient) ListListenerCertificatesWithOptions(request *aliga.ListListenerCertificatesRequest, runtime *util.RuntimeOptions) (_result *aliga.ListListenerCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListenerCertificates"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &aliga.ListListenerCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *GaClient) ListListenerCertificates(request *aliga.ListListenerCertificatesRequest) (_result *aliga.ListListenerCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &aliga.ListListenerCertificatesResponse{}
	_body, _err := client.ListListenerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *GaClient) UpdateAdditionalCertificateWithListenerWithOptions(request *aliga.UpdateAdditionalCertificateWithListenerRequest, runtime *util.RuntimeOptions) (_result *aliga.UpdateAdditionalCertificateWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorId)) {
		query["AcceleratorId"] = request.AcceleratorId
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAdditionalCertificateWithListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &aliga.UpdateAdditionalCertificateWithListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *GaClient) UpdateAdditionalCertificateWithListener(request *aliga.UpdateAdditionalCertificateWithListenerRequest) (_result *aliga.UpdateAdditionalCertificateWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &aliga.UpdateAdditionalCertificateWithListenerResponse{}
	_body, _err := client.UpdateAdditionalCertificateWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *GaClient) UpdateListenerWithOptions(request *aliga.UpdateListenerRequest, runtime *util.RuntimeOptions) (_result *aliga.UpdateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !tea.BoolValue(util.IsUnset(request.BackendPorts)) {
		query["BackendPorts"] = request.BackendPorts
	}

	if !tea.BoolValue(util.IsUnset(request.Certificates)) {
		query["Certificates"] = request.Certificates
	}

	if !tea.BoolValue(util.IsUnset(request.ClientAffinity)) {
		query["ClientAffinity"] = request.ClientAffinity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.HttpVersion)) {
		query["HttpVersion"] = request.HttpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortRanges)) {
		query["PortRanges"] = request.PortRanges
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		query["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyProtocol)) {
		query["ProxyProtocol"] = request.ProxyProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestTimeout)) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		query["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.XForwardedForConfig)) {
		query["XForwardedForConfig"] = request.XForwardedForConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListener"),
		Version:     tea.String("2019-11-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &aliga.UpdateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *GaClient) UpdateListener(request *aliga.UpdateListenerRequest) (_result *aliga.UpdateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &aliga.UpdateListenerResponse{}
	_body, _err := client.UpdateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
