package unicloud

type sdkResponse interface {
	GetSuccess() bool
	GetErrorCode() string
	GetErrorMessage() string
	GetReturnCode() int
	GetReturnDesc() string
}

type sdkResponseBase struct {
	Success *bool              `json:"success,omitempty"`
	Header  *map[string]string `json:"header,omitempty"`
	Error   *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
	ReturnCode *int    `json:"ret,omitempty"`
	ReturnDesc *string `json:"desc,omitempty"`
}

func (r *sdkResponseBase) GetReturnCode() int {
	if r.ReturnCode == nil {
		return 0
	}

	return *r.ReturnCode
}

func (r *sdkResponseBase) GetReturnDesc() string {
	if r.ReturnDesc == nil {
		return ""
	}

	return *r.ReturnDesc
}

func (r *sdkResponseBase) GetSuccess() bool {
	if r.Success == nil {
		return false
	}

	return *r.Success
}

func (r *sdkResponseBase) GetErrorCode() string {
	if r.Error == nil {
		return ""
	}

	return r.Error.Code
}

func (r *sdkResponseBase) GetErrorMessage() string {
	if r.Error == nil {
		return ""
	}

	return r.Error.Message
}

var _ sdkResponse = (*sdkResponseBase)(nil)
