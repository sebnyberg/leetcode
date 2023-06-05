package p1755closestsubsequencesum

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minAbsDifference(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		goal int
		want int
	}{
		{[]int{2886, 2141, 6046, -9031, 5378, -8576, 109}, -512655773, 512638166},
		{[]int{5, -7, 3, 5}, 6, 0},
		{[]int{7, -9, 15, -2}, -5, 1},
		{[]int{1, 2, 3}, -7, 7},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minAbsDifference(tc.nums, tc.goal))
		})
	}
}

func minAbsDifference(nums []int, goal int) int {
	// Divide nums in half and compute all possible sums on each side.
	n := len(nums)
	left := []int{0}
	for _, x := range nums[:n/2] {
		m := len(left)
		for _, y := range left[:m] {
			left = append(left, x+y)
		}
	}
	right := []int{0}
	for _, x := range nums[n/2:] {
		m := len(right)
		for _, y := range right[:m] {
			right = append(right, x+y)
		}
	}
	sort.Ints(left)
	sort.Ints(right)
	res := math.MaxInt64
	for _, x := range left {
		rest := goal - x
		j := sort.SearchInts(right, rest)
		if j < len(right) {
			res = min(res, abs(rest-right[j]))
		}
		if j >= 1 {
			res = min(res, abs(rest-right[j-1]))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
