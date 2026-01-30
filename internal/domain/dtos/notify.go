package dtos

type NotifyTestPushReq struct {
	Provider string `json:"provider"`
	AccessId string `json:"accessId"`
}

type NotifyTestPushResp struct{}
