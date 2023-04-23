package p1416restorethearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfArrays(t *testing.T) {
	for i, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"1000", 10000, 1},
		{"1000", 10, 0},
		{"1317", 2000, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfArrays(tc.s, tc.k))
		})
	}
}

const mod = 1e9 + 7

func numberOfArrays(s string, k int) int {
	// Since k <= 10^9, there can at most be 9 digits in a number.
	// For each position, we can calculate the number of ways that we can end in
	// that position. For this we need to keep track of the number of ways we
	// could have ended in the prior i-9 positions, but lets do all for now.
	n := len(s)
	// dp[i] = number of valid ways to partition the string s[:i]
	dp := make([]int, n+1)
	m := len(fmt.Sprint(k))
	ok := func(t string) bool {
		if t[0] == '0' {
			return false
		}
		var x int
		for i := range t {
			c := int(t[i] - '0')
			x = (x * 10) + c
		}
		return x <= k
	}
	dp[0] = 1
	for i := range s {
		for j := max(0, i-m); j <= i; j++ {
			if !ok(s[j : i+1]) {
				continue
			}
			dp[i+1] = (dp[i+1] + dp[j]) % mod
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
