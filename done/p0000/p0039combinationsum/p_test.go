package p0039combinationsum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum(t *testing.T) {
	for _, tc := range []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.candidates, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, combinationSum(tc.candidates, tc.target))
		})
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	nums := make([]int, 0, len(candidates))
	res := make([][]int, 0)
	findCombinationSums(candidates, target, nums, &res)
	return res
}

func findCombinationSums(candidates []int, target int, nums []int, res *[][]int) {
	switch {
	case target < 0:
		return
	case target == 0:
		a := make([]int, len(nums))
		copy(a, nums)
		*res = append(*res, a)
		return
	}

	for i, n := range candidates {
		nums = append(nums, n)
		findCombinationSums(candidates[i:], target-n, nums, res)
		nums = nums[:len(nums)-1]
	}
}
