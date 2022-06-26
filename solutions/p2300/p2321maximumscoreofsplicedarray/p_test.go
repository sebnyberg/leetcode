package p2321maximumscoreofsplicedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSplicedArray(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{60, 60, 60}, []int{10, 90, 10}, 210},
		{[]int{20, 40, 20, 70, 30}, []int{50, 20, 50, 40, 20}, 220},
		{[]int{7, 11, 13}, []int{1, 1, 1}, 31},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, maximumsSplicedArray(tc.nums1, tc.nums2))
		})
	}
}

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	// Assuming that we swap from one to the other,
	// We can calculate the increase in value by adding each
	// corresponding value from one to the other, then using
	// whateverthefucks algorithm to keep adding values while the sum
	// of values is larger than zero (I think).

	// So the problem becomes:
	// Calculate sums of nums1 and nums2
	// Calculate delta of nums1-nums2
	// Find the subarray with the largest and smallest sum
	// The largest sum can be added to nums2s sum
	// -1 * smallest sum can be added to nums1s sum

	sums := [2]int{}
	n := len(nums1)
	var delta [2][]int
	delta[0] = make([]int, n)
	delta[1] = make([]int, n)
	for i := range nums1 {
		sums[0] += nums1[i]
		sums[1] += nums2[i]
		delta[0][i] = nums2[i] - nums1[i]
		delta[1][i] = nums1[i] - nums2[i]
	}
	largestSequenceSum := func(a []int) int {
		var sum int
		var maxSum int
		for _, x := range a {
			sum += x
			maxSum = max(maxSum, sum)
			if sum < 0 {
				sum = 0
			}
		}
		return maxSum
	}
	d1 := largestSequenceSum(delta[0])
	d2 := largestSequenceSum(delta[1])
	res1 := sums[0] + d1
	res2 := sums[1] + d2
	return max(res1, res2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
