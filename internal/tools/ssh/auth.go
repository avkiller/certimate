package ssh

type AuthMethodType string

const (
	AuthMethodTypeNone     AuthMethodType = "none"
	AuthMethodTypePassword AuthMethodType = "password"
	AuthMethodTypeKey      AuthMethodType = "key"
)
