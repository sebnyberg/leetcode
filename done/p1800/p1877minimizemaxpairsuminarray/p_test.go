package p1877minimizemaxpairsuminarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPairSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 5, 2, 3}, 7},
		{[]int{3, 5, 4, 2, 4, 6}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minPairSum(tc.nums))
		})
	}
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	var maxSum int
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		maxSum = max(maxSum, nums[r]+nums[l])
		l++
		r--
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
