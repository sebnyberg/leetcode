package p1955countnumberofspecialsubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSpecialSubsequences(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 2, 0, 1, 2}, 7},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{2, 2, 0, 0}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countSpecialSubsequences(tc.nums))
		})
	}
}

const mod = 1e9 + 7

func countSpecialSubsequences(nums []int) int {
	var dp [3]int
	for _, n := range nums {
		switch n {
		case 0:
			dp[0] = dp[0]*2 + 1
		case 1:
			dp[1] = dp[1]*2 + dp[0]
		case 2:
			dp[2] = dp[2]*2 + dp[1]
		}
		dp[n] %= mod
	}
	return dp[2]
}
