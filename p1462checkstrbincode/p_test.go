package p1462checkstrbincode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasAllCodes(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want bool
	}{
		{"00110110", 2, true},
		{"00110", 2, true},
		{"0110", 1, true},
		{"0110", 2, false},
		{"0000000001011100", 4, false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, hasAllCodes(tc.s, tc.k))
		})
	}
}
func hasAllCodes(s string, k int) bool {
	nelements := 1 << k
	if len(s) != k*nelements {
		return false
	}
	found := make([]bool, nelements)
	for i := 0; i < len(s); i += k {
		v := 0
		for j := 0; j < k; j++ {
			v <<= 1
			v += int(s[i+k-1-j] - '0')
		}
		found[v] = true
	}
	for _, v := range found {
		if !v {
			return false
		}
	}
	return true
}
