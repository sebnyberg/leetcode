package p0962maximumwidthramp

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxWidthRamp(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 1, 2}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxWidthRamp(tc.nums))
		})
	}
}

func maxWidthRamp(nums []int) int {
	stack := []int{}
	var res int
	for i := range nums {
		j := sort.Search(len(stack), func(k int) bool {
			return nums[stack[k]] <= nums[i]
		})
		if j == len(stack) {
			// No smaller number exists
			stack = append(stack, i)
		} else {
			res = max(res, i-stack[j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
