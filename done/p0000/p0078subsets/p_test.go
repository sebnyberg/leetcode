package p0078subsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subsets(tc.nums))
		})
	}
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	collectSubsets(0, []int{}, nums, &res)
	return res
}

func collectSubsets(idx int, prefix []int, nums []int, result *[][]int) {
	if idx == len(nums) {
		*result = append(*result, prefix)
		return
	}
	// Option 1: add the current number
	prefixCpy := make([]int, len(prefix))
	copy(prefixCpy, prefix)
	prefixCpy = append(prefixCpy, nums[idx])
	collectSubsets(idx+1, prefixCpy, nums, result)

	// Option 2: do nothing
	collectSubsets(idx+1, prefix, nums, result)
}
