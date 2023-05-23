package leetcode

import (
	"fmt"
)

type SegTree[T any] struct {
	tree  []T
	n     int
	m     int
	agg   func(a, b T) T
	empty T
}

// Init initializes the segment tree with the provided values. The default value
// is used for padding and must retain the required invariant given by the agg
// function. For example, if agg is min, then defaultVal should be
// math.MaxInt64, if agg is sum, then defaultVal should be 0, and so on.
func (t *SegTree[T]) Init(nums []T, agg func(a, b T) T, defaultVal T) {
	n := 1
	t.m = len(nums)
	t.empty = defaultVal
	for n < t.m {
		n *= 2
	}
	t.n = n

	// Bunch of boilerplate to effeciently resize the tree to match the size of
	// the input.
	t.tree = t.tree[:cap(t.tree)]
	r := min(len(t.tree), n*2)
	for i := 0; i < r; i++ {
		t.tree[i] = t.empty
	}
	if len(t.tree) < n*2 {
		t.tree = append(t.tree, make([]T, n*2-len(t.tree))...)
	}
	t.Reset()

	copy(t.tree[n:], nums)
	for i := n - 1; i >= 1; i-- {
		t.tree[i] = agg(t.tree[i*2], t.tree[i*2+1])
	}

	t.agg = agg
}

func (t *SegTree[T]) Reset() {
	for i := 1; i < t.n*2; i++ {
		t.tree[i] = t.empty
	}
}

// Update the value of a single element.
// TODO needs update when update range is implemented, current implementation
// overwrites pending lazy changes. Once update-range is in place, it is better
// to just call UpdateRange for the interval [i,i]
func (t *SegTree[T]) Update(i int, v T) {
	t.tree[i+t.n] = v
	for j := (i + t.n) / 2; j >= 1; j /= 2 {
		t.tree[j] = t.agg(t.tree[j*2], t.tree[j*2+1])
	}
}

// UpdateRange updates the range [lo,hi] to have the value T
// Currently not supported.
func (t *SegTree[T]) UpdateRange(lo, hi int, v T) {
	// TODO: implement, and it should be lazy
	panic("todo")
}

// i = index in tree slice
// qlo = queried low end
// qhi = queried high end
// lo = low-end of range covered by this section of the segtree (tree[i])
// hi = high-end of range [...]
func (t *SegTree[T]) query(i, qlo, qhi, lo, hi int) T {
	if qhi < lo || qlo > hi {
		return t.empty
	}
	if qlo <= lo && qhi >= hi {
		// This entire section of the segtree is covered by the query range.
		// This is exactly where the segtree shines!
		return t.tree[i]
	}
	mid := lo + (hi-lo)/2
	left := t.query(i*2, qlo, qhi, lo, mid)
	right := t.query(i*2+1, qlo, qhi, mid+1, hi)
	res := t.agg(left, right)
	return res
}

// QueryRange finds the aggregated value for the range [qlo,qhi].
func (t *SegTree[T]) QueryRange(qlo, qhi int) T {
	if qhi >= t.m {
		s := fmt.Sprintf("index %v out of bounds, slice has len %v", qhi, t.m)
		panic(s)
	}
	if qlo < 0 {
		s := fmt.Sprintf("index %v out of bounds", qlo)
		panic(s)
	}
	return t.query(1, qlo, qhi, 0, t.n-1)
}

// QueryValRange finds the start and end index such that the range includes the
// values [qlo,qhi]. This assumes that the tree is ordered, and that a .less
// function exists (?).
// TODO: not implemented
func (t *SegTree[T]) QueryValRange(qlo, qhi T) [2]int {
	panic("todo")
}

// Search finds the lowest index in the tree for which the aggregated value >=
// val, or the length of the input data if no such index was found.
func (t *SegTree[T]) Search(val T) int {
	panic("todo")
}
