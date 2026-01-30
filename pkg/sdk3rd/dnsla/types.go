package dnsla

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

type DomainRecord struct {
	Id            string `json:"id"`
	GroupId       string `json:"groupId"`
	GroupName     string `json:"groupName"`
	Domain        string `json:"domain"`
	DisplayDomain string `json:"displayDomain"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
}

type DnsRecord struct {
	Id          string `json:"id"`
	DomainId    string `json:"domainId"`
	GroupId     string `json:"groupId"`
	GroupName   string `json:"groupName"`
	LineId      string `json:"lineId"`
	LineCode    string `json:"lineCode"`
	LineName    string `json:"lineName"`
	Type        int32  `json:"type"`
	Host        string `json:"host"`
	DisplayHost string `json:"displayHost"`
	Data        string `json:"data"`
	DisplayData string `json:"displayData"`
	Ttl         int32  `json:"ttl"`
	Weight      int32  `json:"weight"`
	Preference  int32  `json:"preference"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}
