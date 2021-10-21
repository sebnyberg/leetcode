package p1493longestsubarrayof1safterdeletingoneelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 1}, 2},
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 0, 0, 1, 1, 1, 0, 1}, 4},
		{[]int{0, 0, 0}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, longestSubarray(tc.nums))
		})
	}
}

func longestSubarray(nums []int) int {
	var res int
	prev, cur := 0, 0
	var didDelete bool
	for _, n := range nums {
		if n == 1 {
			cur++
			continue
		}
		didDelete = true
		res = max(res, prev+cur)
		prev, cur = cur, 0
	}
	if didDelete {
		res = max(res, prev+cur)
	} else {
		res = max(res, cur-1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
