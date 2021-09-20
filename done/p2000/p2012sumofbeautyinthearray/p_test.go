package p2012sumofbeautiyinthearray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumOfBeauties(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{2, 4, 6, 4}, 1},
		{[]int{3, 2, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, sumOfBeauties(tc.nums))
		})
	}
}

func sumOfBeauties(nums []int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < len(nums); i++ {
		left[i] = max(left[i-1], nums[i-1])
	}
	right[n-1] = math.MaxInt32
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = min(right[i+1], nums[i+1])
	}
	var beauty int
	for i := 1; i < len(nums)-1; i++ {
		// condition 1
		if left[i] < nums[i] && nums[i] < right[i] {
			beauty += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			beauty++
		}
	}
	return beauty
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
