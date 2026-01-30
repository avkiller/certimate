package apisix

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type SslUpdateRequest = SslCertificate

type SslUpdateResponse = SslCertificate

func (c *Client) SslUpdate(sslId string, req *SslUpdateRequest) (*SslUpdateResponse, error) {
	return c.SslUpdateWithContext(context.Background(), sslId, req)
}

func (c *Client) SslUpdateWithContext(ctx context.Context, sslId string, req *SslUpdateRequest) (*SslUpdateResponse, error) {
	httpreq, err := c.newRequest(http.MethodPut, fmt.Sprintf("/ssls/%s", url.PathEscape(sslId)))
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &SslUpdateResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
