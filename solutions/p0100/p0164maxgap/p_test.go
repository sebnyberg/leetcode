package p0164maxgap

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumGap(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1, 1, 1, 5, 5, 5, 5, 5}, 4},
		{[]int{1, 3, 100}, 97},
		{[]int{10}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumGap(tc.nums))
		})
	}
}

func maximumGap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Find minimum and maximum values
	// These are used to find the smallest possible max distance
	// between two numbers in nums
	minVal, maxVal := nums[0], nums[0]
	for _, n := range nums {
		minVal = min(minVal, n)
		maxVal = max(maxVal, n)
	}

	if maxVal == minVal {
		return 0
	}

	// The smallest maximum distance is caused by a series where the numbers
	// between each pair of numbers is exactly the same, e.g. 0, 10, 20
	// This gives a minimum max gap of 20-0 / 2 => 10
	gap := max(1, (maxVal-minVal)/(len(nums)-1)-1)

	// Add one extra bucket here due to flooring in the integer division
	nbuckets := 1 + (maxVal-minVal)/gap
	bucketMin := make([]int, nbuckets)
	bucketMax := make([]int, nbuckets)
	for i := range bucketMin {
		bucketMin[i] = math.MaxInt32
		bucketMax[i] = math.MinInt32
	}

	// Fill buckets with values from nums
	for _, n := range nums {
		idx := (n - minVal) / gap
		bucketMin[idx] = min(bucketMin[idx], n)
		bucketMax[idx] = max(bucketMax[idx], n)
	}

	prevMax := bucketMax[0]
	maxGap := bucketMax[0] - bucketMin[0]
	for i := 1; i < len(bucketMin); i++ {
		if bucketMin[i] == math.MaxInt32 {
			// empty bucket
			continue
		}
		maxGap = max(maxGap, bucketMin[i]-prevMax)
		prevMax = bucketMax[i]
	}

	return maxGap
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
