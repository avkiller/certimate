package ratpanel

type sdkResponse interface {
	GetMessage() string
}

type sdkResponseBase struct {
	Message *string `json:"msg,omitempty"`
}

func (r *sdkResponseBase) GetMessage() string {
	if r.Message == nil {
		return ""
	}

	return *r.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)
