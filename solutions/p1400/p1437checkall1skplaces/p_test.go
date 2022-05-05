package p1437checkall1skplaces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kLengthApart(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 0, 0, 1, 0, 1}, 2, false},
		{[]int{1, 1, 1, 1, 1}, 0, true},
		{[]int{1, 1, 1, 1, 1}, 1, false},
		{[]int{0, 1, 0, 1}, 1, true},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.nums, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, kLengthApart(tc.nums, tc.k))
		})
	}
}

func kLengthApart(nums []int, k int) bool {
	c := k
	for _, n := range nums {
		if n != 1 {
			c++
			continue
		}
		if c < k {
			return false
		}
		c = 0
	}
	return true
}
