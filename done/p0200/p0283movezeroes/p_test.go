package p0283movezeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_moveZeroes(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			moveZeroes(tc.nums)
			require.Equal(t, tc.want, tc.nums)
		})
	}
}

func moveZeroes(nums []int) {
	n := len(nums)
	l := 0
	r := 1
	for r < n && l < n-1 {
		for l < n-1 && nums[l] != 0 {
			l++
		}
		r = l
		for r < n && nums[r] == 0 {
			r++
		}
		if l < n-1 && r < n {
			nums[r], nums[l] = nums[l], nums[r]
		}
	}
}
