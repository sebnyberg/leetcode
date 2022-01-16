package p2134minimumswapstogroupall1stogether2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwaps(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 0, 1, 1, 0, 0}, 1},
		{[]int{0, 1, 1, 1, 0, 0, 1, 1, 0}, 2},
		{[]int{1, 1, 0, 0, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minSwaps(tc.nums))
		})
	}
}

func minSwaps(nums []int) int {
	// This exercise is easy once you understand the problem.
	// The goal is to fill a window with ones (circular or no)
	// The length of the window is the count of ones in the input.
	// The optimal solution is the window which has the least amount of zeroes
	n := len(nums)
	var width int
	for _, num := range nums {
		if num == 1 {
			width++
		}
	}
	var onesCount int
	for i := n - width; i < n; i++ {
		if nums[i] == 1 {
			onesCount++
		}
	}
	minSwap := math.MaxInt32
	for i, v := range nums {
		if v == 1 {
			onesCount++
		}
		if nums[(n+i-width)%n] == 1 {
			onesCount--
		}
		minSwap = min(minSwap, width-onesCount)
	}
	return minSwap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
