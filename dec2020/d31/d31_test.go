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
	largestArea := 0
	for i, height := range heights {
		start, end := i, i
		for ; start >= 0 && heights[start] >= height; start-- {
		}
		for ; end < len(heights) && heights[end] >= height; end++ {
		}
		largestArea = max(largestArea, height*(end-start-1))
	}
	return largestArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
