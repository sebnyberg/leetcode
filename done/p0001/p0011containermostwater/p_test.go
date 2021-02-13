package p0011containermostwater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxArea(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, maxArea(tc.in))
		})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) (maxWater int) {
	l, r := 0, len(height)-1
	for l != r {
		if height[l] < height[r] {
			maxWater = max(maxWater, height[l]*(r-l))
			l++
			continue
		}
		maxWater = max(maxWater, height[r]*(r-l))
		r--
	}
	return maxWater
}
