package p0468validateipaddress

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validIPAddress(t *testing.T) {
	for _, tc := range []struct {
		queryIP string
		want    string
	}{
		{"172.16.254.1", "IPv4"},
		{"2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"256.256.256.256", "Neither"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.queryIP), func(t *testing.T) {
			require.Equal(t, tc.want, validIPAddress(tc.queryIP))
		})
	}
}

func validIPAddress(queryIP string) string {
	validIPv4Part := func(s string) bool {
		if len(s) == 0 || len(s) > 3 {
			return false
		}
		if s[0] == '0' {
			return len(s) == 1
		}
		v, err := strconv.Atoi(s)
		return err == nil && v >= 0 && v <= 255
	}
	hex := "0123456789abcdefABCDEF"
	validIPv6Part := func(s string) bool {
		if len(s) == 0 || len(s) > 4 {
			return false
		}
		for _, ch := range s {
			if !strings.ContainsRune(hex, ch) {
				return false
			}
		}
		return true
	}
	parts := strings.Split(queryIP, ".")
	if len(parts) == 4 {
		for _, p := range parts {
			if !validIPv4Part(p) {
				return "Neither"
			}
		}
		return "IPv4"
	}
	parts = strings.Split(queryIP, ":")
	if len(parts) == 8 {
		for _, p := range parts {
			if !validIPv6Part(p) {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
