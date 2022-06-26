package p2311longestbinarysubsequencelessthanorequaltok

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"1001010", 5, 5},
		{"00101001", 1, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestSubsequence(tc.s, tc.k))
		})
	}
}

func longestSubsequence(s string, k int) int {
	var l int
	var val int
	var maxDist int
	var onesRemoved int
	for r := 0; r < len(s); r++ {
		val <<= 1
		val += int(byte(s[r] - '0'))
		for ; l != r && val > k; l++ {
			if s[l] == '0' {
				continue
			}
			onesRemoved++
			pow := 1 << (r - l)
			val -= pow
		}
		maxDist = max(maxDist, r-onesRemoved+1)
	}
	return maxDist
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
