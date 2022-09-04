package p2401longestnicesubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestNiceSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 8, 48, 10}, 3},
		{[]int{3, 1, 5, 11, 13}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, longestNiceSubarray(tc.nums))
		})
	}
}

func longestNiceSubarray(nums []int) int {
	var res, l, bits int
	for i, x := range nums {
		for l != i && bits&x != 0 {
			bits ^= nums[l]
			l++
		}
		bits |= nums[i]
		res = max(res, i-l+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
