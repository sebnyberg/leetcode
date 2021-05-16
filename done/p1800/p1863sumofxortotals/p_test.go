package p1863sumofxortotals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsetXORSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3}, 6},
		{[]int{5, 1, 6}, 28},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subsetXORSum(tc.nums))
		})
	}
}

func subsetXORSum(nums []int) int {
	res := make([]int, 0)
	helper(nums, 0, 0, &res)
	var val int
	for _, v := range res {
		val += v
	}
	return val
}

func helper(nums []int, idx, cur int, res *[]int) {
	if idx == len(nums) {
		*res = append(*res, cur)
		return
	}
	helper(nums, idx+1, cur, res)
	cur ^= nums[idx]
	helper(nums, idx+1, cur, res)
}
