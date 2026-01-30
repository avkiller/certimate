package tls

import (
	"crypto/tls"
)

// 创建并返回一个兼容低版的 [tls.Config] 对象。
//
// 出参：
//   - config: [tls.Config] 对象。
func NewCompatibleConfig() *tls.Config {
	var suiteIds []uint16
	for _, suite := range tls.CipherSuites() {
		suiteIds = append(suiteIds, suite.ID)
	}
	for _, suite := range tls.InsecureCipherSuites() {
		suiteIds = append(suiteIds, suite.ID)
	}

	return &tls.Config{
		MinVersion:   tls.VersionTLS10,
		CipherSuites: suiteIds,
	}
}

// 创建并返回一个不安全的 [tls.Config] 对象。
//
// 出参：
//   - config: [tls.Config] 对象。
func NewInsecureConfig() *tls.Config {
	config := NewCompatibleConfig()
	config.InsecureSkipVerify = true
	return config
}
