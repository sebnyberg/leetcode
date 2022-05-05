package p2025maximizenumberofwaystopartitionarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_waysToPartition(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, -1, 2}, 3, 1},
		{[]int{0, 0, 0}, 1, 2},
		{[]int{22, 4, -25, -20, -15, 15, -16, 7, 19, -10, 0, -13, -14}, -33, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, waysToPartition(tc.nums, tc.k))
		})
	}
}

func waysToPartition(nums []int, k int) int {
	// Calculate sum difference when changing a number in nums to k
	// Calculate pre-sum from right to left
	// Iterate from left to right, keeping track of the sum
	// Take presum - left sum. If the difference exists as the result of changing
	// a number on the left side, then add the number of partitions for that
	// change.
	// This process has covered all partitions for a change such that the change
	// was on the left side of the pivot.
	// Reverse the list of numbers.
	// Re-do the calculation once again - this accounts for the number of equal
	// partitions such that the change was on the right side of the partition.
	// Invert nums and re-do calculation
	n := len(nums)
	validLeftPartitions, noChangeCount := countValidLeftPartitions(nums, k)
	// reverse
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	// Count partitions where the change is on the right side
	validRightPartitions, noChangeCount2 := countValidLeftPartitions(nums, k)
	if noChangeCount != noChangeCount2 {
		panic("wat")
	}
	for i := range validRightPartitions {
		validLeftPartitions[i] += validRightPartitions[n-i-1]
	}
	res := noChangeCount
	for _, count := range validLeftPartitions {
		res = max(res, count)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// countValidLeftPartitions counts the number of valid partitions in nums such
// that one number in the left partition was changed to k for an index in nums.
// Returns the number of such valid partitions, and the number of valid
// partitions given no change at all.
func countValidLeftPartitions(nums []int, k int) ([]int, int) {
	n := len(nums)
	validPartitions := make([]int, n)
	var noChangeCount int
	diffIndices := make(map[int][]int)
	var rightSum int
	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
	}
	// Add zeroth index
	if nums[0] != k {
		diffIndices[k-nums[0]] = append(diffIndices[nums[0]-k], 0)
	}
	leftSum := nums[0]
	for i := 1; i < n; i++ {
		if rightSum == leftSum {
			noChangeCount++
		} else {
			if indices, exists := diffIndices[rightSum-leftSum]; exists {
				for _, idx := range indices {
					validPartitions[idx]++
				}
			}
		}
		leftSum += nums[i]
		rightSum -= nums[i]
		if d := k - nums[i]; d != 0 {
			diffIndices[d] = append(diffIndices[d], i)
		}
	}
	return validPartitions, noChangeCount
}
