package p1793maximizescoreofagoodsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumScore(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{5, 5, 4, 5, 4, 1, 1, 1}, 0, 20},
		{[]int{1, 4, 3, 7, 4, 5}, 3, 15},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumScore(tc.nums, tc.k))
		})
	}
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	minVal := nums[k]
	maxVal := nums[k]
	l, r := k, k
	for l > 0 || r < n-1 {
		switch {
		case l == 0:
			r++
		case r == n-1:
			l--
		case nums[l-1] < nums[r+1]:
			r++
		default:
			l--
		}
		minVal = min(minVal, min(nums[r], nums[l]))
		maxVal = max(maxVal, minVal*(r-l+1))
	}
	return maxVal
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
