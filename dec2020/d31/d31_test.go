package d31_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestRectangleArea(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 1, 2}, 3},
		{[]int{2147483647, 0, 2147483647, 0, 2147483647, 0, 2147483647, 0, 2147483647, 0}, 2147483647},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, largestRectangleArea(tc.in))
		})
	}
}

func largestRectangleArea(heights []int) int {
	// Keep a stack of indices for the heights
	// E.g. [0,1,2,3] => heights[1,6,8,9]
	stack := make([]int, 0)

	// When a height is larger or equal to the height of the last index
	// in the stack, pop the stack and calculate the height of each pop.
	heights = append(heights, 0)

	maxArea := 0

	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stackHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			var stackWidth int
			if len(stack) == 0 {
				stackWidth = i
			} else {
				stackWidth = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, stackHeight*stackWidth)
		}
		stack = append(stack, i)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
