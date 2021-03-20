package p0189rotatearr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotate(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			rotate(tc.nums, tc.k)
			require.Equal(t, tc.want, tc.nums)
		})
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	rev(nums[:n-k])
	rev(nums[n-k:])
	rev(nums)
}

func rev(ns []int) {
	for i, j := 0, len(ns)-1; i < j; i, j = i+1, j-1 {
		ns[i], ns[j] = ns[j], ns[i]
	}
}
