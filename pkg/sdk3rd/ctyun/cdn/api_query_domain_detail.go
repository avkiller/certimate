package cdn

import (
	"context"
	"net/http"

	qs "github.com/google/go-querystring/query"
)

type QueryDomainDetailRequest struct {
	Domain        *string `json:"domain,omitempty"         url:"domain,omitempty"`
	ProductCode   *string `json:"product_code,omitempty"   url:"product_code,omitempty"`
	FunctionNames *string `json:"function_names,omitempty" url:"function_names,omitempty"`
}

type QueryDomainDetailResponse struct {
	sdkResponseBase

	ReturnObj *DomainDetail `json:"returnObj,omitempty"`
}

func (c *Client) QueryDomainDetail(req *QueryDomainDetailRequest) (*QueryDomainDetailResponse, error) {
	return c.QueryDomainDetailWithContext(context.Background(), req)
}

func (c *Client) QueryDomainDetailWithContext(ctx context.Context, req *QueryDomainDetailRequest) (*QueryDomainDetailResponse, error) {
	httpreq, err := c.newRequest(http.MethodGet, "/v1/domain/query-domain-detail")
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

	result := &QueryDomainDetailResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
