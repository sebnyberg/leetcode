package p0581shortestunsortedcontinuousarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findUnsortedSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 3, 3}, 2},
		{[]int{1, 3, 2, 2, 2}, 4},
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findUnsortedSubarray(tc.nums))
		})
	}
}

func findUnsortedSubarray(nums []int) int {
	stack := make([]int, 0, len(nums))
	start := math.MaxInt32
	end := 0
	maxVal := math.MinInt32
	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > num {
			// Pop from stack
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				start = min(start, stack[len(stack)-1])
			} else {
				start = -1
			}
			end = max(end, i)
		}
		if num < maxVal {
			end = i
		}
		maxVal = max(maxVal, num)
		stack = append(stack, i)
	}
	if start == math.MaxInt32 {
		return 0
	}

	return end - start
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
