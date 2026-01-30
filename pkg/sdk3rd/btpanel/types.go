package btpanel

type sdkResponse interface {
	GetStatus() *bool
	GetMessage() *string
}

type sdkResponseBase struct {
	Status  *bool   `json:"status,omitempty"`
	Message *string `json:"msg,omitempty"`
}

func (r *sdkResponseBase) GetStatus() *bool {
	return r.Status
}

func (r *sdkResponseBase) GetMessage() *string {
	return r.Message
}
