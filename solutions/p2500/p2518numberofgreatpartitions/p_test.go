package p2518numberofgreatpartitions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPartitions(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{
			[]int{96, 40, 22, 98, 9, 97, 45, 22, 79, 57, 95, 62},
			505,
			0,
		},
		{[]int{3, 3, 3}, 4, 0},
		{[]int{6, 6}, 2, 2},
		{[]int{1, 2, 3, 4}, 4, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countPartitions(tc.nums, tc.k))
		})
	}
}

const mod = 1e9 + 7

func countPartitions(nums []int, k int) int {
	// dp[i] = number of ways to get sum i, where i <= k
	dp := make([]int, k+1)
	dp[0] = 1
	var totalSum int
	nsums := 1
	for _, x := range nums {
		nsums = (nsums * 2) % mod
		totalSum += x
		for y := k; y >= x; y-- {
			dp[y] = (dp[y-x] + dp[y]) % mod
		}
	}
	if totalSum < k*2 {
		// Note: this heuristic is not correct.
		// It does not adequately check whether a solution actually exists.
		// For example, if k = 10, totalSum = 21, then this would not be
		// triggered. But there may still be no valid solutions, and the mod
		// below could return a large value e.g. 8183891.
		//
		// The problem is that it is impossible to check whether the number of
		// small elements is larger than the total number of elements (due to
		// the mod). Another (invalid) heuristic would be to keep track of the
		// power of the number of small vs all sums but that would also fail on
		// boundaries (or grow too large to handle in an int).
		//
		// I assume that the only real solution would be to use a big int.
		return 0
	}
	var nsmall int
	for x := 0; x < k; x++ {
		nsmall = (nsmall + dp[x]) % mod
	}
	res := (nsums - (nsmall * 2) + 2*mod) % mod
	return res
}
