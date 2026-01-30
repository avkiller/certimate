package dns

import (
	"context"
	"net/http"
)

type CreateRecordRequest struct {
	ZoneName   *string                      `json:"zone_name,omitempty"`
	DomainName *string                      `json:"domain_name,omitempty"`
	ViewId     *int32                       `json:"view_id,omitempty"`
	Type       *string                      `json:"type,omitempty"`
	Records    []*CreateRecordRequestRecord `json:"record,omitempty"`
	Ttl        *int32                       `json:"ttl,omitempty"`
	Mode       *int32                       `json:"mode,omitempty"`
	AutoMerge  *int32                       `json:"auto_merge,omitempty"`
}

type CreateRecordRequestRecord struct {
	Values []*CreateRecordRequestRecordValue `json:"values,omitempty"`
	Weight *int32                            `json:"weight,omitempty"`
}

type CreateRecordRequestRecordValue struct {
	Value  *string `json:"value,omitempty"`
	Status *int32  `json:"status,omitempty"`
}

type CreateRecordResponse struct {
	sdkResponseBase
	DomainName     *string                       `json:"domain_name,omitempty"`
	DomainRecordId *int64                        `json:"domain_record_id,omitempty"`
	ViewId         *int64                        `json:"view_id,omitempty"`
	Records        []*CreateRecordResponseRecord `json:"records,omitempty"`
}

type CreateRecordResponseRecord struct {
	GroupId     *int64                             `json:"group_id,omitempty"`
	GroupStatus *int32                             `json:"group_status,omitempty"`
	Values      []*CreateRecordResponseRecordValue `json:"value,omitempty"`
	Weight      *int32                             `json:"weight,omitempty"`
}

type CreateRecordResponseRecordValue struct {
	ValueId *int64  `json:"id,omitempty"`
	Value   *string `json:"value,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

func (c *Client) CreateRecord(req *CreateRecordRequest) (*CreateRecordResponse, error) {
	return c.CreateRecordWithContext(context.Background(), req)
}

func (c *Client) CreateRecordWithContext(ctx context.Context, req *CreateRecordRequest) (*CreateRecordResponse, error) {
	httpreq, err := c.newRequest(http.MethodPost, "/v1/record/")
	if err != nil {
		return nil, err
	} else {
		httpreq.SetBody(req)
		httpreq.SetContext(ctx)
	}

	result := &CreateRecordResponse{}
	if _, err := c.doRequestWithResult(httpreq, result); err != nil {
		return result, err
	}

	return result, nil
}
