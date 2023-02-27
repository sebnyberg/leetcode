package p2569handlingsumqueriesafterupdate

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_handleQuery(t *testing.T) {
	for i, tc := range []struct {
		nums1   []int
		nums2   []int
		queries [][]int
		want    []int64
	}{
		{[]int{1, 0, 1}, []int{0, 0, 0}, [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}}, []int64{3}},
		{[]int{1}, []int{5}, [][]int{{2, 0, 0}, {3, 0, 0}}, []int64{5}},
		{
			[]int{0, 0, 0, 0, 1, 0, 1, 1, 1},
			[]int{35, 29, 21, 34, 8, 48, 22, 43, 37},
			leetcode.ParseMatrix("[[1,4,7],[3,0,0],[2,27,0],[3,0,0],[1,0,3],[3,0,0],[2,6,0],[1,3,8],[2,13,0],[3,0,0],[3,0,0],[3,0,0],[2,2,0],[2,28,0],[3,0,0],[3,0,0],[2,25,0],[3,0,0],[3,0,0],[1,2,5]]"),
			[]int64{277, 331, 331, 445, 445, 445, 625, 625, 775, 775},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, handleQuery(tc.nums1, tc.nums2, tc.queries))
		})
	}
}

func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	// There are three types of "queries"
	//
	// 1. Flip 0s to 1s and vice versa in nums1 within a certain interval
	// 2. Total sum += p * sum(nums1)
	// 3. Return total sum
	//
	// One thing we know is that any range is either inverted or not. If it is
	// inverted, then the count is inverted as well. That is, if the regular
	// count of 1s in [3, 9] is 3, then the inverted count must be (9-3+1)-3 = 4
	// This means that the original count only needs to be stored once. After
	// that, we only need to keep track of inversions.
	//
	// The difficult part is to flip all bits in a range without going through
	// each individual element one by one. What if we use a segment tree to
	// store flips?
	//
	// I will not go through how segment trees work here, but the idea is that
	// each segment in nums1 will keep a tally of ones and zeroes. When
	// flipping, the two are swapped.
	//
	// Why not use left/right + count of ones? Because we need to fill the
	// segment tree to be a power-of-two, so some entries are not actually part
	// of nums1.
	//

	// Build the tree
	n := 1
	m := len(nums1)
	for n < m {
		n *= 2
	}
	t := make([]int, n*2)
	lazyFlips := make([]int, n*2)
	for i := range nums1 {
		t[n+i] = nums1[i]
	}
	for i := n - 1; i >= 1; i-- {
		t[i] = t[i*2] + t[i*2+1]
	}

	// push consumes any delayed flips and propagate them further down the tree.
	push := func(i, l, r int) {
		if lazyFlips[i] == 0 {
			return
		}
		if lazyFlips[i]&1 > 0 {
			t[i] = (r - l + 1) - t[i]
		}
		if l != r {
			lazyFlips[i*2] += lazyFlips[i]
			lazyFlips[i*2+1] += lazyFlips[i]
		}
		lazyFlips[i] = 0
		return
	}

	// update flips a range of bits
	var update func(i, tl, tr, l, r int) int
	update = func(i, tl, tr, l, r int) int {
		// perform any outstanding flips and propagate
		push(i, tl, min(tr, m-1))

		if r < tl || l > tr {
			return t[i] // no flip
		}

		if tl >= l && tr <= r {
			lazyFlips[i]++
			push(i, tl, min(tr, m-1))
			return t[i]
		}

		// range is within this segment, but may not cover all children.
		mid := tl + (tr-tl)/2
		a := update(2*i, tl, mid, l, r)
		b := update(2*i+1, mid+1, tr, l, r)
		t[i] = a + b
		return t[i]
	}

	var res []int64
	var sum int
	for _, x := range nums2 {
		sum += x
	}
	for _, q := range queries {
		switch q[0] {
		case 1:
			update(1, 0, n-1, q[1], q[2])
		case 2:
			sum += t[1] * q[1]
		case 3:
			res = append(res, int64(sum))
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
