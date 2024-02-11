package p2926maximumbalancedsubsequencesum

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxBalancedSubsequenceSum(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int64
	}{
		// {[]int{3, 3, 5, 6}, 14},
		// {[]int{5, -1, -3, 8}, 13},
		{[]int{-2, -1}, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxBalancedSubsequenceSum(tc.nums))
		})
	}
}

func maxBalancedSubsequenceSum(nums []int) int64 {
	// This is a classical multi-indexing problem.
	//
	// The first idea here is that whether or not a prior value can be regarded
	// for a given number in nums depends on the difference in delta between
	// that prior value and the current number.
	//
	// In other words, given i and j, then nums[j] can succeed nums[i] iff
	// nums[j]-j >= nums[i]-i
	//
	// If we consider numbers in order by (a) position and (b) delta, then we
	// can be sure than any subsequence that precedes the currently evaluated
	// number is a viable prior subsequence as well. In order to find the
	// maximum sum subsequence prior to the current position, we can use a
	// segment tree.
	//
	m := len(nums)
	n := 1
	for n < m {
		n <<= 1
	}
	t := make([]int, n*2)
	for i := range t {
		t[i] = math.MinInt32
	}
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		ii := idx[i]
		jj := idx[j]
		di := nums[ii] - ii
		dj := nums[jj] - jj
		if di == dj {
			return ii < jj
		}
		return di < dj
	})

	var query func(i, tl, tr, l, r int) int
	query = func(i, tl, tr, l, r int) int {
		if r < tl || l > tr {
			return math.MinInt32
		}

		if tl >= l && tr <= r {
			// Whole tree segment covered by [l,r]
			return t[i]
		}
		mid := tl + (tr-tl)/2
		a := query(i*2, tl, mid, l, r)
		b := query(i*2+1, mid+1, tr, l, r)
		return max(a, b)
	}

	res := math.MinInt32
	for _, i := range idx {
		// Find most promising prior subsequence that is valid
		x := query(1, 0, n-1, 0, i)
		y := max(nums[i], nums[i]+x)

		res = max(res, y)

		// Update the tree
		t[n+i] = y
		for j := (n + i) / 2; j >= 1; j /= 2 {
			t[j] = max(t[j*2], t[j*2+1])
		}
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
