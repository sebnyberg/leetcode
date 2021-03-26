package p0213houserobber2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rob(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{0}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, rob(tc.nums))
		})
	}
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var withFirst [2]int
	var withoutFirst [2]int
	for i, n := range nums {
		withFirst[1], withFirst[0] = max(withFirst[1], withFirst[0]+n), withFirst[1]
		if i > 0 {
			withoutFirst[1], withoutFirst[0] = max(withoutFirst[1], withoutFirst[0]+n), withoutFirst[1]
		}
	}
	return max(withFirst[0], withoutFirst[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
