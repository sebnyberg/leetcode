package p1004maxconsecutiveones3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestOnes(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{0, 0, 0, 1}, 4, 4},
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, longestOnes(tc.nums, tc.k))
		})
	}
}

func longestOnes(nums []int, k int) int {
	n := len(nums)
	l, r := 0, 0
	var zeroes int
	var res int
	for {
		// Read until all 'k' have been used up
		for r < n && zeroes <= k {
			if nums[r] == 0 {
				zeroes++
			}
			r++
		}
		res = max(res, r-l-1)
		if r == n {
			break
		}
		// Move left cursor until k is OK
		for l < r && k < zeroes {
			if nums[l] == 0 {
				zeroes--
			}
			l++
		}
	}
	if zeroes <= k {
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
