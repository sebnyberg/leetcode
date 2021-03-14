package p0152maxprodsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProduct(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxProduct(tc.nums))
		})
	}
}

func maxProduct(nums []int) int {
	prevMax, prevMin := nums[0], nums[0]
	totalMax := nums[0]
	for _, n := range nums[1:] {
		a, b := prevMax*n, prevMin*n
		prevMax = max(n, max(a, b))
		prevMin = min(n, min(a, b))
		if prevMax > totalMax {
			totalMax = prevMax
		}
	}
	return totalMax
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
