package p0046permutations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permuteUnique(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, permuteUnique(tc.nums))
		})
	}
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	findPerms(nums, []int{}, &res)
	return res
}

func findPerms(nums []int, prefix []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, prefix)
	}

	// For each position, create a prefix for each unique number
	// found in nums, remove the number from nums, and pass it
	// onto the child function
	seen := make(map[int]struct{}, len(nums))
	for i, n := range nums {
		if _, exists := seen[n]; exists {
			continue
		}
		seen[n] = struct{}{}
		prefixCpy := make([]int, len(prefix)+1)
		copy(prefixCpy, prefix)
		prefixCpy[len(prefix)] = n

		nCpy := make([]int, len(nums))
		copy(nCpy, nums)
		nCpy = append(nCpy[:i], nCpy[i+1:]...)

		findPerms(nCpy, prefixCpy, res)
	}
}
