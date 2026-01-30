package rainyun

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

type SslRecord struct {
	ID         int64  `json:"ID"`
	UID        int64  `json:"UID"`
	Domain     string `json:"Domain"`
	Issuer     string `json:"Issuer"`
	StartDate  int64  `json:"StartDate"`
	ExpireDate int64  `json:"ExpDate"`
	UploadTime int64  `json:"UploadTime"`
}

type SslDetail struct {
	Cert       string `json:"Cert"`
	Key        string `json:"Key"`
	Domain     string `json:"DomainName"`
	Issuer     string `json:"Issuer"`
	StartDate  int64  `json:"StartDate"`
	ExpireDate int64  `json:"ExpDate"`
	RemainDays int64  `json:"RemainDays"`
}
