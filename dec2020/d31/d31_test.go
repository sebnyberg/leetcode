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
		{[]int{2147483647, 0, 2147483647, 0, 2147483647, 0, 2147483647, 0, 2147483647, 0}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, largestRectangleArea(tc.in))
		})
	}
}

func largestRectangleArea(heights []int) int {
	largestArea := 0
	for i := 0; i < len(heights); i++ {
		if i > 1 && heights[i] < heights[i-1] {
			continue
		}
		height := heights[i]
		for step := 1; step <= height; step++ {
			j := i + 1
			for ; j < len(heights) && heights[j] >= step; j++ {
			}
			largestArea = max(largestArea, step*(j-i))
		}
	}
	return largestArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
