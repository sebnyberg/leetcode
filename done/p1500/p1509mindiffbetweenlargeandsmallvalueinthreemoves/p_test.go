package p1509mindiffbetweenlargeandsmallvalueinthreemoves

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 5, 0, 10, 14}, 1},
		{[]int{6, 6, 0, 1, 1, 4, 6}, 2},
		{[]int{1, 5, 6, 14, 15}, 1},
		{[]int{5, 3, 2, 4}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minDifference(tc.nums))
		})
	}
}

func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	// Skip all largest numbers
	minVal := nums[0]
	maxVal := nums[n-4]
	res := maxVal - minVal
	for i := 1; i <= 3; i++ {
		minVal = nums[i]
		maxVal = nums[n-4+i]
		res = min(res, maxVal-minVal)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
