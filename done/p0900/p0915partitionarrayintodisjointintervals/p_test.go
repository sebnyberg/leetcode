package p0915partitionarrayintodisjointintervals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partitionDisjoint(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{5, 0, 3, 8, 6}, 3},
		{[]int{1, 1, 1, 0, 6, 12}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, partitionDisjoint(tc.nums))
		})
	}
}

func partitionDisjoint(nums []int) int {
	// This problem can be reformulated as:
	// Partition nums such that the smallest number from the right partition is
	// greater than or equal to the greatest number from the left partition.
	//
	minRight := make([]int, len(nums))
	n := len(nums)
	minRight[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		minRight[i] = min(minRight[i+1], nums[i])
	}
	maxLeft := nums[0]
	for i := 1; i < n; i++ {
		if maxLeft <= minRight[i] {
			return i
		}
		maxLeft = max(maxLeft, nums[i])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
