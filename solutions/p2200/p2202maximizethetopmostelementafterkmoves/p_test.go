package p2202maximizethetopmostelementafterkmoves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumTop(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{91, 98},
			2, 91,
		},
		{[]int{35, 43, 23, 86, 23, 45, 84, 2, 18, 83, 79, 28, 54, 81, 12, 94, 14, 0, 0, 29, 94, 12, 13, 1, 48, 85, 22, 95, 24, 5, 73, 10, 96, 97, 72, 41, 52, 1, 91, 3, 20, 22, 41, 98, 70, 20, 52, 48, 91, 84, 16, 30, 27, 35, 69, 33, 67, 18, 4, 53, 86, 78, 26, 83, 13, 96, 29, 15, 34, 80, 16, 49},
			15, 94,
		},
		{[]int{5, 2, 2, 4, 0, 6}, 4, 5},
		{[]int{2}, 1, -1},
		{[]int{2}, 2, 2},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumTop(tc.nums, tc.k))
		})
	}
}

func maximumTop(nums []int, k int) int {
	if len(nums) == 1 {
		if k%2 == 0 {
			return nums[0]
		} else {
			return -1
		}
	}

	var maxVal int
	for i := 0; i < min(k-1, len(nums)); i++ {
		maxVal = max(maxVal, nums[i])
	}
	if k < len(nums) {
		maxVal = max(maxVal, nums[k])
	}

	return maxVal
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
