package p2790maximumnumberofgroupswithincreasinglength

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxIncreasingGroups(t *testing.T) {
	for i, tc := range []struct {
		usageLimits []int
		want        int
	}{
		{[]int{6, 3, 4, 1, 1, 4, 3, 1, 2, 3, 5}, 7},
		{[]int{2, 3}, 2},
		{[]int{1, 1, 5}, 2},
		{[]int{2, 2, 2}, 3},
		{[]int{1, 2, 5}, 3},
		{[]int{2, 1, 2}, 2},
		{[]int{1, 1}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxIncreasingGroups(tc.usageLimits))
		})
	}
}

func maxIncreasingGroups(usageLimits []int) int {
	// This seems to be quite a tricky problem.
	//
	// Consider the "pyramid" given by [0, 4, 2, 2, 1, 1]
	//
	// 1
	// 1
	// 1 2 3
	// 1 2 3 4 5
	//
	// This pyramid can form four groups. This is because the excess width of
	// the bottom level can be redistributed to the top level. It seems to me
	// that given a certain condition, a sufficiently clever algorithm should be
	// able to distribute the numbers in such a way that the result is
	// constrained by the total count being >= the arithmetic sum
	// 1 + 2 + ... + n = n(n-1)/2
	//
	// So one constraining factor is simply the total limit of all numbers.
	// Another one is just the number of distinct values.
	//
	// That is to say: that if we have [10, 10, 10], then the answer is clearly
	// 3. However, if we have [1, 1, 1, 1], the answer is 2. The first is
	// constrained by the variety of numbers, and the other by the numbers'
	// cardinality. I wonder if there is a middle-ground where both matter.
	//
	// Ah there is one case, let's call it "skimming". When two numbers have a
	// cardinality higher than the number of groups, then one of them cannot
	// contribute to the total sum of limits used to distribute numbers as a
	// pyramid. For example:
	//
	// 1 2
	// 1 2
	// 1 2
	// 1 2 3 4
	//
	sort.Sort(sort.Reverse(sort.IntSlice(usageLimits)))

	n := len(usageLimits)
	ok := func(m int) bool {
		// Count sum of limits beyond m (if applicable)
		// Then create an ascending sum from the rightmost element to the first
		// one. If the carry ever becomes negative, then there is no solution
		//
		if m > len(usageLimits) {
			return false
		}
		var carry int
		for i := m; i < n; i++ {
			carry += usageLimits[i]
		}
		for i := m - 1; i >= 0; i-- {
			want := m - i
			if usageLimits[i] < want {
				carry -= want - usageLimits[i]
				if carry < 0 {
					return false
				}
			} else {
				carry += usageLimits[i] - want
			}
		}
		return true
	}

	lo, hi := 1, math.MaxInt64/10
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !ok(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
