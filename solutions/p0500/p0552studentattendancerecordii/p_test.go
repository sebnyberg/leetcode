package p0552studentattendancerecordii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkRecord(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{4, 43},
		{3, 19},
		{2, 8},
		{1, 3},
		{10101, 183236316},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, checkRecord(tc.n))
		})
	}
}

const mod = 1e9 + 7

func checkRecord(n int) int {
	// dp[i][j] = i absent days so far, j last days late
	var dp [2][3]int
	dp[1][0] = 1 // 'A'
	dp[0][1] = 1 // 'L'
	dp[0][0] = 1 // 'P'

	for i := 1; i < n; i++ {
		var next [2][3]int

		// Addin 'P' to 'PPPP' is fine.
		// Adding 'P' to anything ending with a sequence of 'L' is also fine
		next[0][0] = (dp[0][0] + dp[0][1] + dp[0][2]) % mod

		// We can add 'P' to any existing records with one 'A'
		next[1][0] = (dp[1][0] + dp[1][1] + dp[1][2]) % mod
		// And we can add 'A' to any records with no 'A'
		next[1][0] = (next[1][0] + dp[0][0] + dp[0][1] + dp[0][2]) % mod

		// Adding an 'L' is fine for all groups that don't end with two L's
		next[0][1] = dp[0][0]
		next[1][1] = dp[1][0]
		next[0][2] = dp[0][1]
		next[1][2] = dp[1][1]

		dp = next
	}

	var res int
	for i := range dp {
		for j := range dp[i] {
			res = (res + dp[i][j]) % mod
		}
	}
	return res
}
