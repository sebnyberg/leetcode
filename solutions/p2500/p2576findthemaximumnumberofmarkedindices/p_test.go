package p2576findthemaximumnumberofmarkedindices

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNumOfMarkedIndices(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{9, 10, 10, 14, 15, 16, 17, 17, 24, 28, 29, 31, 32, 40, 42, 48, 51, 55, 64, 68, 71, 83, 98, 99, 99, 100}, 26},
		{[]int{9, 2, 5, 4}, 4},
		{[]int{7, 6, 8}, 0},
		{[]int{3, 5, 2, 4}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxNumOfMarkedIndices(tc.nums))
		})
	}
}

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	r := len(nums) / 2
	var count int
	var l int
	for l < len(nums)/2 {
		for r < len(nums) && nums[r] < nums[l]*2 {
			r++
		}
		if r == len(nums) {
			break
		}
		count += 2
		r++
		l++
	}
	return count
}
