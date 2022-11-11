package p0956tallestbillboard

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_tallestBillboard(t *testing.T) {
	for i, tc := range []struct {
		rods []int
		want int
	}{
		{[]int{1, 2, 3, 6}, 6},
		{[]int{1, 2, 3, 4, 5, 6}, 10},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, tallestBillboard(tc.rods))
		})
	}
}

func tallestBillboard(rods []int) int {
	// As usual for hard problems, enumeration of combinations and partitioning
	// does not work. You must find a trick that reduces the state to it's most
	// critical information. For this problem, that critical information is a
	// combination of the current shared length of beams and difference between
	// beams.
	//
	// The idea is to, for each beam, consider all possible differences between
	// beam heights. Then do:
	//
	// 1. Nothing
	// 2. Add current to shorter beam
	// 3. Add current to higher beam
	//
	// For case (2), we increase the minimum height.
	//
	// The result can be retrieved from dp[0]
	var sum int
	for _, r := range rods {
		sum += r
	}
	dp := make([]int, sum/2+1)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = 0
	prev := make([]int, sum/2+1)
	for _, r := range rods {
		for i := range prev {
			prev[i] = dp[i]
		}
		for d := 0; d <= sum/2; d++ {
			// Add rod to smaller beam, increasing sum by d and changing delta
			// to abs(d-r)
			dd := abs(d - r)
			if dd <= sum/2 {
				dp[dd] = max(dp[dd], prev[d]+min(r, d))
			}

			// Add rod to larger beam
			dd = d + r
			if dd <= sum/2 {
				dp[dd] = max(dp[dd], prev[d])
			}
		}
	}
	return dp[0]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
