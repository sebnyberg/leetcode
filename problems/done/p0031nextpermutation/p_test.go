package p0031nextpermutation

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextPermuation(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{1}, []int{1}},
		{[]int{1, 5, 1}, []int{5, 1, 1}},
		{[]int{2, 3, 0, 2, 4, 1}, []int{2, 3, 0, 4, 1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			nextPermutation(tc.in)
			require.Equal(t, tc.want, tc.in)
		})
	}
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			// Find smallest increase above nums[i-1]
			minDiff := nums[i] - nums[i-1]
			minDiffIdx := i
			for j := i + 1; j < len(nums); j++ {
				diff := nums[j] - nums[i-1]
				if diff > 0 && diff < minDiff {
					minDiff = diff
					minDiffIdx = j
				}
			}
			nums[i-1], nums[minDiffIdx] = nums[minDiffIdx], nums[i-1]
			sort.Ints(nums[i:])
			return
		}
	}
	sort.Ints(nums)
}
