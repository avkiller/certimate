package hostname

import (
	"crypto/x509"
	"net"
	"strings"
)

// 检查目标主机名是否匹配待匹配主机名。
//
// 入参：
//   - match: 待匹配主机名。可以是泛域名，如 "*.example.com"。
//   - candidate: 目标主机名。如 "sub.example.com"。
//
// 出参：
//   - 是否匹配。
func IsMatch(match, candidate string) bool {
	if match == "" || candidate == "" {
		return false
	}

	mockCert := &x509.Certificate{}
	if ip := net.ParseIP(match); ip != nil {
		mockCert.IPAddresses = []net.IP{ip}
	} else {
		if strings.EqualFold(match, candidate) {
			return true
		}
		mockCert.DNSNames = []string{match}
	}
	return mockCert.VerifyHostname(candidate) == nil
}
