package p0643maximumaveragesubarray

import (
	"fmt"
	"math"
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
	maxSum := math.MinInt32
	for i, val := range nums {
		sum += val
		if i >= k {
			sum -= nums[i-k]
		}
		if i >= k-1 {
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return float64(maxSum) / float64(k)
}
