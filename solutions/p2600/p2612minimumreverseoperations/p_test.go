package p2612minimumreverseoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reachableRange(t *testing.T) {
	for i, tc := range []struct {
		n         int
		i         int
		k         int
		wantStart int
		wantEnd   int
	}{
		{9, 2, 7, 4, 8},
		{9, 0, 2, 1, 1},
		{9, 8, 2, 7, 7},
		{9, 6, 2, 5, 7},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			a, b := reachableRange(tc.n, tc.i, tc.k)
			require.Equal(t, tc.wantStart, a, "start")
			require.Equal(t, tc.wantEnd, b, "end")
		})
	}
}

func Test_minReverseOperations(t *testing.T) {
	for i, tc := range []struct {
		n      int
		p      int
		banned []int
		k      int
		want   []int
	}{
		{5, 0, []int{}, 2, []int{0, 1, 2, 3, 4}},
		{4, 2, []int{0, 1, 3}, 1, []int{-1, -1, 0, -1}},
		{4, 0, []int{1, 2}, 4, []int{0, -1, -1, 1}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minReverseOperations(tc.n, tc.p, tc.banned, tc.k))
		})
	}
}

func minReverseOperations(n int, p int, banned []int, k int) []int {
	// This problem is quite hard.
	//
	// First, it is worth noting that we want to work with odd and even
	// positions in the array separately. This is because an odd k can only
	// reverse 1 into positions at an even distance from the current position.
	// Similarly, with an even k, the 1 can only be moved into an odd distance
	// from the current position.
	//
	// By splitting the array into odd and even positions, the reachable
	// positions can be described as a range rather than single positions,
	// making it possible to do a range update with a segment tree.
	//
	m := 1
	neven := n/2 + n&1
	for m < neven {
		m *= 2
	}
	var segtree [2][]bool
	segtree[0] = make([]bool, m*2)
	segtree[1] = make([]bool, m*2)

	var res [2][]int
	res[0] = make([]int, neven)
	res[1] = make([]int, n/2)
	for i := range res {
		for j := range res[i] {
			res[i][j] = -2
		}
	}

	// mark marks the provided index with the provided value
	mark := func(segtree []bool, res []int, i, val int) {
		segtree[m+i] = true
		for k := (m + i) / 2; k >= 1; k /= 2 {
			segtree[k] = segtree[k*2] && segtree[k*2+1]
			if !segtree[k] {
				break
			}
		}
		if i < len(res) {
			res[i] = val
		}
	}

	for _, i := range banned {
		mark(segtree[i&1], res[i&1], i/2, -1)
	}

	// Mark out-of-bounds elements as seen
	for i := neven; i < m; i++ {
		mark(segtree[0], res[0], i, -1)
	}
	for i := n / 2; i < m; i++ {
		mark(segtree[1], res[1], i, -1)
	}

	type pos struct {
		i   int
		odd int
	}

	// Update marks the range
	// [lo,hi] is the current range in the segment tree
	// [qlo,qhi] is the range being updated by the caller
	// j is 0 if using even segtree, otherwise 1
	var update func(segtree []bool, res []int, next *[]pos, i, lo, hi, qlo, qhi, val, odd int)
	update = func(segtree []bool, res []int, next *[]pos, i, lo, hi, qlo, qhi, val, odd int) {
		if qhi < lo || qlo > hi {
			// Skip
			return
		}

		if lo >= qlo && hi <= qhi {
			// This range of the segment tree should be marked as "done"
			if segtree[i] {
				// If it is already done, return
				return
			}

			// Mark all nodes as done, adding unseen nodes to next iteration
			for j := lo; j <= hi; j++ {
				if res[j] != -2 {
					continue
				}
				res[j] = val

				mark(segtree, res, j, val)
				*next = append(*next, pos{j, odd})
			}
			return
		}

		// Split query into left/right
		mid := lo + (hi-lo)/2
		update(segtree, res, next, i*2, lo, mid, qlo, qhi, val, odd)
		update(segtree, res, next, i*2+1, mid+1, hi, qlo, qhi, val, odd)
	}

	curr := []pos{}
	next := []pos{}

	// Start by marking p as done
	// This will populate next with the first node.
	update(segtree[p&1], res[p&1], &curr, 1, 0, m-1, p/2, p/2, 0, p&1)

	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			ii := x.i*2 + x.odd
			l, r := reachableRange(n, ii, k)
			update(segtree[l&1], res[l&1], &next, 1, 0, m-1, l/2, r/2, steps, l&1)
		}
		curr, next = next, curr
	}

	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = res[i&1][i/2]
	}
	for i := range ret {
		if ret[i] == -2 {
			ret[i] = -1
		}
	}
	return ret
}

func reachableRange(n, i, k int) (int, int) {
	leftMost := i - k + 1
	rightMost := i + k - 1
	if leftMost < 0 {
		d := -leftMost
		leftMost += d * 2
	}
	if rightMost >= n {
		d := rightMost - n + 1
		rightMost -= d * 2
	}
	return leftMost, rightMost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
