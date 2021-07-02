package p0643maximumaveragesubarray1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxAverage(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxAverage(tc.nums, tc.k))
		})
	}
}

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	n := len(nums)
	for j := k; j < n; j++ {
		sum -= nums[j-k]
		sum += nums[j]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
