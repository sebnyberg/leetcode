package p0075sortcolors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortColors(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			sortColors(tc.nums)
			require.Equal(t, tc.want, tc.nums)
		})
	}
}

func sortColors(nums []int) {
	var r, g, b int
	for _, n := range nums {
		switch n {
		case 0:
			r++
		case 1:
			g++
		case 2:
			b++
		}
	}

	var pos int
	for ; pos < r; pos++ {
		nums[pos] = 0
	}
	for ; pos < r+g; pos++ {
		nums[pos] = 1
	}
	for ; pos < r+g+b; pos++ {
		nums[pos] = 2
	}
}
