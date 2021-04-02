package p0482licencekeyformatting

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_licenceKeyFormatting(t *testing.T) {
	for _, tc := range []struct {
		S    string
		K    int
		want string
	}{
		{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"2-5g-3-J", 2, "2-5G-3J"},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.S, tc.K), func(t *testing.T) {
			require.Equal(t, tc.want, licenseKeyFormatting(tc.S, tc.K))
		})
	}
}

func licenseKeyFormatting(S string, K int) string {
	// Assumption: leave the first group full unless there is less than
	stripped := strings.ToUpper(strings.Join(strings.Split(S, "-"), ""))
	ngroups := ((len(stripped) - 1) / K) + 1
	firstGroupSize := ((len(stripped) - 1) % K) + 1
	res := []byte(stripped[:firstGroupSize])
	stripped = stripped[firstGroupSize:]
	for i := 1; i < ngroups; i++ {
		res = append(res, '-')
		res = append(res, stripped[:K]...)
		stripped = stripped[K:]
	}
	return string(res)
}
