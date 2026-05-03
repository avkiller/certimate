package hostname

import (
	"crypto/x509"
	"net"
	"strings"

	xcert "github.com/certimate-go/certimate/pkg/utils/cert"
)

// 检查目标主机名是否匹配待匹配主机名。
// 兼容目标主机名开头是 "." 的情况（视为泛域名）。
//
// 入参：
//   - pattern: 待匹配主机名，可以是泛域名。如 "*.example.com"。
//   - hostname: 目标主机名。如 "sub.example.com"。
//
// 出参：
//   - 是否匹配。
func IsMatch(pattern, hostname string) bool {
	if pattern == "" || hostname == "" {
		return false
	}

	mockCert := &x509.Certificate{}
	if ip := net.ParseIP(pattern); ip != nil {
		mockCert.IPAddresses = []net.IP{ip}
	} else {
		if strings.EqualFold(pattern, hostname) {
			return true
		}
		mockCert.DNSNames = []string{pattern}
	}
	return IsMatchByCertificate(mockCert, hostname)
}

// 检查目标主机名是否匹配证书。
// 兼容目标主机名开头是 "." 的情况（视为泛域名）。
//
// 入参：
//   - certPEM: 证书 PEM 内容。
//   - hostname: 目标主机名。如 "sub.example.com"。
//
// 出参：
//   - 是否匹配。
func IsMatchByCertificatePEM(certPEM string, hostname string) bool {
	if certPEM == "" || hostname == "" {
		return false
	}

	certX509, err := xcert.ParseCertificateFromPEM(certPEM)
	if err != nil {
		return false
	}

	return IsMatchByCertificate(certX509, hostname)
}

// 检查目标主机名是否匹配证书。
// 兼容目标主机名开头是 "." 的情况（视为泛域名）。
//
// 入参：
//   - certX509: 证书 X509 对象。
//   - hostname: 目标主机名。如 "sub.example.com"。
//
// 出参：
//   - 是否匹配。
func IsMatchByCertificate(certX509 *x509.Certificate, hostname string) bool {
	if certX509 == nil || hostname == "" {
		return false
	}

	if strings.HasPrefix(hostname, "*.") || strings.HasPrefix(hostname, ".") {
		for _, dn := range certX509.DNSNames {
			if strings.EqualFold(strings.TrimPrefix(dn, "*"), strings.TrimPrefix(hostname, "*")) {
				return true
			}
		}
	}

	return certX509.VerifyHostname(hostname) == nil
}
