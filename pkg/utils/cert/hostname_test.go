package cert_test

import (
	"testing"

	xcert "github.com/certimate-go/certimate/pkg/utils/cert"
)

func TestCertUtil_Hostname(t *testing.T) {
	t.Run("MatchHostname", func(t *testing.T) {
		testCases := []struct {
			wildcard string
			target   string
			expected bool
		}{
			{"*.example.com", "sub.example.com", true},
			{"*.example.com", "sub.sub.example.com", false},
			{"*.example.com", "example.com", false},

			{"*.*.example.com", "a.b.example.com", false},
			{"*.*.example.com", "a.example.com", false},
			{"*.*.example.com", "a.b.c.example.com", false},

			{"example.com", "example.com", true},
			{"example.com", "wrong.com", false},

			{"", "example.com", false},
			{"*.example.com", "", false},

			{"*.sub.example.com", "a.sub.example.com", true},
			{"*.sub.example.com", "a.b.sub.example.com", false},
			{"*.sub.example.com", "sub.example.com", false},

			{"*.Example.COM", "sub.example.com", true},
			{"*.EXAMPLE.COM", "SUB.EXAMPLE.COM", true},
		}

		for _, tc := range testCases {
			result := xcert.MatchHostname(tc.wildcard, tc.target)
			status := "✓"
			pf := t.Logf
			if result != tc.expected {
				status = "✗"
				pf = t.Errorf
			}

			pf("%s Wildcard: %-20s Target: %-20s Expected: %-5v Got: %-5v\n", status, tc.wildcard, tc.target, tc.expected, result)
		}
	})
}
