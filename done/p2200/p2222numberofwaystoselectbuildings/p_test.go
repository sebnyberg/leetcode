package p2222numberofwaystoselectbuildings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfWays(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int64
	}{
		{"001101", 6},
		{"11100", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfWays(tc.s))
		})
	}
}

func numberOfWays(s string) int64 {
	var dp [4][2]int
	dp[0][0] = 1
	dp[0][1] = 1
	for _, ch := range s {
		v := int(ch - '0')
		for i := 3; i >= 1; i-- {
			if v == 0 {
				dp[i][0] += dp[i-1][1]
			} else {
				dp[i][1] += dp[i-1][0]
			}
		}
	}
	return int64(dp[3][0] + dp[3][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
