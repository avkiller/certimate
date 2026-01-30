package internal

import (
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	alidcdn "github.com/alibabacloud-go/dcdn-20180115/v4/client"
	"github.com/alibabacloud-go/tea/dara"
)

// This is a partial copy of https://github.com/alibabacloud-go/dcdn-20180115/blob/master/client/client_context_func.go
// to lightweight the vendor packages in the built binary.
type DcdnClient struct {
	openapi.Client
	DisableSDKError *bool
}

func NewDcdnClient(config *openapiutil.Config) (*DcdnClient, error) {
	client := new(DcdnClient)
	err := client.Init(config)
	return client, err
}

func (client *DcdnClient) Init(config *openapiutil.Config) (_err error) {
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

func (client *DcdnClient) DescribeDcdnUserDomainsWithContext(ctx context.Context, request *alidcdn.DescribeDcdnUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *alidcdn.DescribeDcdnUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.ChangeEndTime) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !dara.IsNil(request.ChangeStartTime) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !dara.IsNil(request.CheckDomainShow) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.WebSiteType) {
		query["WebSiteType"] = request.WebSiteType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alidcdn.DescribeDcdnUserDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *DcdnClient) SetDcdnDomainSSLCertificateWithContext(ctx context.Context, request *alidcdn.SetDcdnDomainSSLCertificateRequest, runtime *dara.RuntimeOptions) (_result *alidcdn.SetDcdnDomainSSLCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}

	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		query["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainSSLCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &alidcdn.SetDcdnDomainSSLCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
