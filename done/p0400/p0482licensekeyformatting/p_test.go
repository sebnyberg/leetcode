package p0482licensekeyformatting

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_licenseKeyFormatting(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"2-5g-3-J", 2, "2-5G-3J"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, licenseKeyFormatting(tc.s, tc.k))
		})
	}
}

func licenseKeyFormatting(s string, k int) string {
	r := strings.NewReplacer("-", "")
	s = r.Replace(s)
	s = strings.ToUpper(s)

	parts := make([]string, 0, len(s)/k)
	firstSize := len(s) % k
	if firstSize > 0 {
		parts = append(parts, s[:firstSize])
	}
	for i := firstSize; i < len(s); i += k {
		parts = append(parts, s[i:i+k])
	}
	return strings.Join(parts, "-")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
