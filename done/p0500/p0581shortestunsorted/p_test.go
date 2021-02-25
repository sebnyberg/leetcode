package p0581shortestunsorted

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findUnsortedSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 2, 2}, 4},
		{[]int{2, 3, 3, 2, 4}, 3},
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 3, 4, 3, 1, 3, 2, 5}, 6},
		{[]int{4, 3, 2, 1, 2}, 5},
		{[]int{4, 3, 2, 1}, 4},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findUnsortedSubarray(tc.nums))
		})
	}
}

func findUnsortedSubarray(nums []int) int {
	// Find start
	var start int
	for ; ; start++ {
		if start == len(nums)-1 {
			return 0
		}
		if nums[start] > nums[start+1] {
			break
		}
	}

	// Find end
	var end int
	for end = len(nums) - 1; ; end-- {
		if end == start {
			panic("wut")
		}
		if nums[end] < nums[end-1] {
			break
		}
	}

	// Find min/end within range to sort
	minVal, maxVal := nums[start], nums[start]
	for i := start; i <= end; i++ {
		minVal = min(minVal, nums[i])
		maxVal = max(maxVal, nums[i])
	}

	start = sort.Search(start, func(i int) bool {
		return nums[i] > minVal
	})

	remainder := len(nums) - 1 - end
	newEndOffset := sort.Search(remainder, func(i int) bool {
		return nums[end+1+i] >= maxVal
	})
	end = end + newEndOffset

	return end - start + 1
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
