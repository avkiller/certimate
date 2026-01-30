package ao

import (
	"context"
	"net/http"

	qs "github.com/google/go-querystring/query"
)

type QueryCertRequest struct {
	Id        *int64  `json:"id,omitempty"         url:"id,omitempty"`
	Name      *string `json:"name,omitempty"       url:"name,omitempty"`
	UsageMode *int32  `json:"usage_mode,omitempty" url:"usage_mode,omitempty"`
}

type QueryCertResponse struct {
	sdkResponseBase

	ReturnObj *struct {
		Result *CertDetail `json:"result,omitempty"`
	} `json:"returnObj,omitempty"`
}

func (c *Client) QueryCert(req *QueryCertRequest) (*QueryCertResponse, error) {
	return c.QueryCertWithContext(context.Background(), req)
}

func (c *Client) QueryCertWithContext(ctx context.Context, req *QueryCertRequest) (*QueryCertResponse, error) {
	httpreq, err := c.newRequest(http.MethodGet, "/ctapi/v1/accessone/cert/query")
	if err != nil {
		return nil, err
	} else {
		values, err := qs.Values(req)
		if err != nil {
			return nil, err
		}

		httpreq.SetQueryParamsFromValues(values)
		httpreq.SetContext(ctx)
	}

	result := &QueryCertResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
