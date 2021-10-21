package p0360sorttransformedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortTransformedArray(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		a, b, c int
		want    []int
	}{
		{[]int{-4, -2, 2, 4}, 0, -1, 5, []int{1, 3, 7, 9}},
		{[]int{-4, -2, 2, 4}, 0, 3, 5, []int{-7, -1, 11, 17}},
		{[]int{-4, -2, 2, 4}, -1, 3, 5, []int{-23, -5, 1, 7}},
		{[]int{-4, -2, 2, 4}, 1, 3, 5, []int{3, 9, 15, 33}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, sortTransformedArray(tc.nums, tc.a, tc.b, tc.c))
		})
	}
}

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	// The formula y=ax^2+bx+c forms a parabola which, depending on the sign of a,
	// will have a min or max value at x = b/2a
	// If a > 0, then numbers from [-inf...b/2a] will be in falling order,
	// and [b/2a...inf] are in asc order.
	// Similarly, for a < 0, numbers [-inf...b/2a] are in ascending order, and
	// [b/2a...inf] are descending.
	// The idea is to find the center point of nums, then merge the ascending
	// portion with the reverse of the other portion.
	n := len(nums)
	transformed := make([]int, n)
	midIdx := n
	for i, x := range nums {
		transformed[i] = a*x*x + b*x + c
		if i > 0 && midIdx == n && (transformed[i]-transformed[i-1])*a > 0 {
			midIdx = i
		}
	}
	left := transformed[:midIdx]
	right := transformed[midIdx:]
	switch {
	case a == 0 && b < 0 || a > 0:
		rev(left)
	case a == 0 && b >= 0 || a < 0:
		rev(right)
	}
	return merge(left, right, n)
}

func merge(left, right []int, n int) []int {
	merged := make([]int, 0, n)
	l, r := 0, 0
	for l+r < n {
		switch {
		case r == len(right) ||
			l < len(left) && left[l] < right[r]:
			merged = append(merged, left[l])
			l++
		default:
			merged = append(merged, right[r])
			r++
		}
	}
	return merged
}

func rev(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[r], nums[l] = nums[l], nums[r]
	}
}
