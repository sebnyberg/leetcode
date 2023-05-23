package leetcode

import (
	"errors"
	"fmt"
)

type SegTree[T any] struct {
	tree  []T
	n     int // number of elements in input
	m     int // bottom level size in the tree
	agg   func(a, b T) T
	empty T
}

// NewSegtree creates a new segment tree. The default value is used for padding
// out-of-bounds values, and must be chosen carefully so that agg(defaultVal,
// val) = val. For example if agg is min(a, b), then defaultVal should be
// math.MaxInt64, if agg is sum(a, b), then defaultVal should be 0, and so on.
func NewSegtree[T any](defaultVal T, agg func(a, b T) T) (*SegTree[T], error) {
	if agg == nil {
		return nil, errors.New("agg is required")
	}
	var t SegTree[T]
	t.agg = agg
	t.empty = defaultVal
	t.m = 1
	return &t, nil
}

// Init fast-inizializes the segtree with the provided values in O(n). If the
// tree is no large enough to accomodate the values, the inner tree is expanded.
// Values outside the range of the provided array are initialized using the
// default value provided to the constructor.
func (t *SegTree[T]) Init(nums []T) {
	t.grow(len(nums))
	copy(t.tree[t.m:], nums)
	for i := t.m + t.n; i < t.m*2; i++ {
		t.tree[i] = t.empty
	}
	for i := t.m - 1; i >= 1; i-- {
		t.tree[i] = t.agg(t.tree[i*2], t.tree[i*2+1])
	}
}

// Reset resets the tree
func (t *SegTree[T]) Reset(n int) {
	t.grow(n)
	for i := 1; i < t.m*2; i++ {
		t.tree[i] = t.empty
	}
}

// grow updates the tree to size n, ensuring that it can be accomodated
func (t *SegTree[T]) grow(n int) {
	if t.agg == nil {
		panic("nil agg function")
	}
	for n > t.m {
		t.m *= 2
	}
	if len(t.tree) < t.m*2 {
		nmissing := t.m*2 - len(t.tree)
		t.tree = append(t.tree, make([]T, nmissing)...)
	}
	t.n = n
}

// Update the value of a single element.
// TODO needs update when update range is implemented, current implementation
// overwrites pending lazy changes. Once update-range is in place, it is better
// to just call UpdateRange for the interval [i,i]
func (t *SegTree[T]) Update(i int, v T) {
	t.tree[i+t.m] = v
	for j := (i + t.m) / 2; j >= 1; j /= 2 {
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
	if qhi >= t.n {
		s := fmt.Sprintf("index %v out of bounds, slice has len %v", qhi, t.n)
		panic(s)
	}
	if qlo < 0 {
		s := fmt.Sprintf("index %v out of bounds", qlo)
		panic(s)
	}
	return t.query(1, qlo, qhi, 0, t.m-1)
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
