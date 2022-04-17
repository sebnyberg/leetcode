package p2243calculatedigitsumofastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_digitSum(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"11111222223", 3, "135"},
		{"00000000", 3, "000"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, digitSum(tc.s, tc.k))
		})
	}
}

func digitSum(s string, k int) string {
	for len(s) > k {
		next := make([]byte, 0, len(s))
		for i := 0; i < len(s); i += k {
			var sum int
			for j := i; j < min(len(s), i+k); j++ {
				sum += int(s[j] - '0')
			}
			a := fmt.Sprint(sum)
			next = append(next, a...)
		}
		s = string(next)
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
