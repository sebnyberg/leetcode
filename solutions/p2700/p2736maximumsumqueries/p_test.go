package p2736maximumsumqueries

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maximumSumQueries(t *testing.T) {
	for i, tc := range []struct {
		nums1   []int
		nums2   []int
		queries [][]int
		want    []int
	}{
		{
			[]int{4, 3, 1, 2},
			[]int{2, 4, 9, 5},
			leetcode.ParseMatrix("[[4,1],[1,3],[2,5]]"),
			[]int{6, 10, 7},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSumQueries(tc.nums1, tc.nums2, tc.queries))
		})
	}
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	// Create a segment tree that contains pairs sorted by nums2 value, and
	// returns the maximum combination of nums2 and nums1 for the given range.
	m := 1
	n := len(nums1)
	nq := len(queries)
	for m < n {
		m <<= 1
	}
	// tree[i] = maximum pair value
	tree := make([]int, m*2)
	for i := 1; i < m*2; i++ {
		tree[i] = 0
	}
	update := func(i, val int) {
		tree[i+m] = val
		for j := (i + m) / 2; j >= 1; j /= 2 {
			tree[j] = max(tree[j*2], tree[j*2+1])
		}
	}
	var query func(i, lo, hi, qlo, qhi int) int
	query = func(i, lo, hi, qlo, qhi int) int {
		if hi < qlo || lo > qhi {
			return 0
		}
		if lo >= qlo && hi <= qhi {
			return tree[i]
		}
		mid := lo + (hi-lo)/2
		l := query(i*2, lo, mid, qlo, qhi)
		r := query(i*2+1, mid+1, hi, qlo, qhi)
		return max(l, r)
	}

	// Initialize pairs
	// i = original index
	// xi = index when sorted by x, ASC
	// yi = index when sorted by y, ASC
	type pair struct {
		i  int
		xi int
		yi int
	}
	pairs := make([]pair, n)
	for i := range pairs {
		pairs[i] = pair{i: i}
	}

	// xidx[i] = j means that nums[j] has the sorted position i
	// For example, if xidx[0] = 9, then nums1[9] is min(nums1[:])
	xidx := make([]int, n)
	yidx := make([]int, n)
	for i := range xidx {
		xidx[i] = i
		yidx[i] = i
	}
	sort.Slice(xidx, func(i, j int) bool {
		a := nums1[xidx[i]]
		b := nums1[xidx[j]]
		return a < b
	})
	sort.Slice(yidx, func(i, j int) bool {
		a := nums2[yidx[i]]
		b := nums2[yidx[j]]
		return a < b
	})
	for i := range xidx {
		// p.xi is the sorted position of the num1
		pairs[xidx[i]].xi = i
		pairs[yidx[i]].yi = i
	}

	// Sort queries by x value ascending
	queryIdx := make([]int, nq)
	for i := range queryIdx {
		queryIdx[i] = i
	}
	sort.Slice(queryIdx, func(i, j int) bool {
		a := queries[queryIdx[i]]
		b := queries[queryIdx[j]]
		return a[0] < b[0]
	})

	// The plan is as follows:
	// Use the segment-tree to store values sorted by Y, ascending
	// Sort queries by X, ascending
	// For each query from high, to low
	// Add all pairs which have an X-value larger than the query x value to the
	// segment tree.
	// Then query the range of valid y-values in the segtree. We know that any
	// entry in the tree is valid in terms of x, because otherwise it would not
	// be in the tree to begin with.
	res := make([]int, nq)
	j := n - 1
	for i := nq - 1; i >= 0; i-- {
		qi := queryIdx[i]
		q := queries[qi]
		x := q[0]
		y := q[1]
		// While there are x's >= q[0], add them to the seg-tree
		for j >= 0 && nums1[xidx[j]] >= x {
			// Add to segtree
			p := pairs[xidx[j]]
			update(p.yi, nums1[p.i]+nums2[p.i])
			j--
		}
		// Find sorted index of y
		// This is the lower-bound of the seg-tree query
		qlo := sort.Search(len(yidx), func(i int) bool {
			return nums2[yidx[i]] >= y
		})
		qres := query(1, 0, m-1, qlo, n-1)
		if qres == 0 {
			res[qi] = -1
			continue
		}
		res[qi] = qres
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
