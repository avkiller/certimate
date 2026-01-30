package dns

type sdkResponse interface {
	GetCode() int
	GetMessage() string
}

type sdkResponseBase struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (r *sdkResponseBase) GetCode() int {
	if r.Code == nil {
		return 0
	}

	return *r.Code
}

func (r *sdkResponseBase) GetMessage() string {
	if r.Message == nil {
		return ""
	}

	return *r.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)

type DnsRecord struct {
	GroupId     *int64            `json:"group_id,omitempty"`
	GroupStatus *int32            `json:"group_status,omitempty"`
	Value       []*DnsRecordValue `json:"value,omitempty"`
	Weight      *int32            `json:"weight,omitempty"`
}

type DnsRecordValue struct {
	Id    *int64  `json:"id,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
	Line  *string `json:"line,omitempty"`
	Ttl   *int32  `json:"ttl,omitempty"`
}
