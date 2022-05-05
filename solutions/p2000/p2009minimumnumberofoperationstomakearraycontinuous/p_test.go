package p2008maximumearningfromtaxi

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{8, 10, 16, 18, 10, 10, 16, 13, 13, 16}, 6},
		{[]int{8, 5, 9, 9, 8, 4}, 2},
		{[]int{4, 2, 5, 3}, 0},
		{[]int{1, 2, 3, 5, 6}, 1},
		{[]int{1, 10, 100, 1000}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums))
		})
	}
}

func minOperations(nums []int) int {
	// It's clear that the worst solution is n-1
	// Idea: sort nums, keep two pointers and a window of max size n-1.
	// The largest possible window meeting the criteria is the optimal choice
	// of numbers such that a continuous array can be formed.
	n := len(nums)
	sort.Ints(nums)
	nums = append(nums, 2e9) // sentinel value
	nums = append(nums, 3e9) // sentinel values
	var l, r int
	// [1, 2, infty]
	nunique := 1
	for r < n && nums[r+1]-nums[l] < n {
		if nums[r+1] != nums[r] {
			nunique++
		}
		r++
	}
	minChanges := n - nunique
	for r < n-1 {
		// move left one step
		l++
		if nums[l] != nums[l-1] {
			nunique--
		}
		// keep moving right until window won't fit
		for r < n && (r < l || nums[r+1]-nums[l] < n) {
			r++
			if nums[r] != nums[r-1] {
				nunique++
			}
		}
		minChanges = min(minChanges, n-nunique)
	}
	return minChanges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
