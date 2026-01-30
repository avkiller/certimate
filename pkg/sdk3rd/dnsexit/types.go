package dnsexit

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
	Type      *string `json:"type,omitempty"`
	Name      *string `json:"name,omitempty"`
	Content   *string `json:"content,omitempty"`
	TTL       *int    `json:"ttl,omitempty"`
	Priority  *int    `json:"priority,omitempty"`
	Overwrite *bool   `json:"overwrite,omitempty"`
}
