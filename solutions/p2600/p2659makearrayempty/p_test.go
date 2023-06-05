package p2659makearrayempty

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countOperationsToEmptyArray(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{1, 2, 4, 3}, 5},
		{[]int{3, 4, -1}, 5},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countOperationsToEmptyArray(tc.nums))
		})
	}
}

func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)

	// segtree boilerplate
	m := 1
	for m < n {
		m <<= 1
	}
	tree := make([]int, m*2)
	for i := m; i < m*2; i++ {
		tree[i] = 1
	}
	for i := m - 1; i >= 1; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}
	var query func(i, lo, hi, qlo, qhi int) int
	query = func(i, lo, hi, qlo, qhi int) int {
		if qhi < lo || qlo > hi {
			return 0
		}
		if lo >= qlo && hi <= qhi {
			return tree[i]
		}
		mid := lo + (hi-lo)/2
		l := query(i*2, lo, mid, qlo, qhi)
		r := query(i*2+1, mid+1, hi, qlo, qhi)
		return l + r
	}
	update := func(j, val int) {
		tree[m+j] = 0
		for k := (m + j) / 2; k >= 1; k /= 2 {
			tree[k] = tree[k*2] + tree[k*2+1]
		}
	}

	// idx[i] = when sorted, number at nums[idx[i]] is at position i
	// i.e. nums[idx[0]] is the minimum value in nums, and so on
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})

	res := idx[0] + 1
	update(idx[0], 0)
	for i := 0; i < len(idx)-1; i++ {
		// Current number is "popped". We want to calculate the distance to move
		// before we can pop the next element.
		// Using the segtree, we can quickly count remaining numbers between the
		// two indices.
		if idx[i+1] < idx[i] {
			// Wrap-around. Calculate distance to end of array, then from start
			r := query(1, 0, m-1, idx[i], n-1)
			l := query(1, 0, m-1, 0, idx[i+1])
			res += r + l
		} else {
			r := query(1, 0, m-1, idx[i], idx[i+1])
			res += r
		}
		update(idx[i+1], 0)
	}

	return int64(res)
}
