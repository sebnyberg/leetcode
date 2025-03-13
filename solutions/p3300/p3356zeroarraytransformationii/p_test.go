package p3356zeroarraytransformationii

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minZeroArray(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		want    int
	}{
		{[]int{2, 0, 2}, leetcode.ParseMatrix("[[0,2,1],[0,2,1],[1,1,3]]"), 2},
		{[]int{4, 3, 2, 1}, leetcode.ParseMatrix("[[1,3,2],[0,2,1]]"), -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minZeroArray(tc.nums, tc.queries))
		})
	}
}

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	delta := make([]int, n)
	ok := func(k int) bool {
		// Check if it is possible to make all elements of nums equal to 0 using at
		// most k operations.
		for i := range delta {
			delta[i] = 0
		}
		for _, q := range queries[:k] {
			l, r, val := q[0], q[1], q[2]
			delta[l] += val
			if r+1 < n {
				delta[r+1] -= val
			}
		}
		var currDelta int
		for i, x := range nums {
			currDelta += delta[i]
			if currDelta < x {
				return false
			}
		}
		return true
	}
	lo, hi := 0, len(queries)+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !ok(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo > len(queries) {
		return -1
	}
	return lo
}
