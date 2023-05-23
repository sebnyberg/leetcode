package p1425constrainedsubsequencesum

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_constrainedSubsetSum(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{-8269, 3217, -4023, -4138, -683, 6455, -3621, 9242, 4015, -3790}, 1, 16091},
		{[]int{10, -2, -10, -5, 20}, 2, 23},
		{[]int{10, 2, -10, 5, 20}, 2, 37},
		{[]int{-1, -2, -3}, 1, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, constrainedSubsetSum(tc.nums, tc.k))
		})
	}
}

func constrainedSubsetSum(nums []int, k int) int {
	// The tricky part of this problem is that it is worth it to include
	// negative values if it can bring two positive subsequences together.
	//
	// We could say that the array nums includes subarrays that contain positive
	// sum subsequences, and we want to determine the minimum cost of combining
	// these subarrays, and whether the cost is warranted
	//
	// So let's split this into subproblems:
	//
	// 1. Detect subarrays that start and end with positive values, and where no
	// two consecutive values violate the constraint given by the problem.
	// 2. Find lowest sum paths between positive sum subarrays
	// 3. Combine subarrays for which the lowest-sum cross-subarray path is
	// small enough to warrant their combination. The maximum sum of a combined
	// subarray is the final result.
	if k == 1 {
		// Handle edge-case of having no positive values at all
		maxVal := math.MinInt32
		for _, x := range nums {
			maxVal = max(maxVal, x)
			if x >= 0 {
				goto hasPositive
			}
		}
		return maxVal
	hasPositive:
	}

	// Skip initial negative values
	var j int
	for nums[j] < 0 {
		j++
	}

	var res int
	var sum int
	for j < len(nums) {
		sum = max(0, sum)

		// Greedily sum until there are no non-negative candidate numbers
		l := j
		for j < len(nums) && (nums[j] >= 0 || j-l < k) {
			if nums[j] >= 0 {
				l = j
				sum += nums[j]
			}
			j++
		}
		res = max(res, sum)

		// Now we are stuck, either because j == len(nums), or because there is
		// a section of negative values blocking us from continuing.
		if j == len(nums) {
			break
		}

		// Find the section of negative values
		for j < len(nums) && nums[j] < 0 {
			j++
		}
		if j == len(nums) {
			break
		}

		// nums[j] >= 0 and the region nums[l+1:j] contains only values smaller
		// than 0. We must find the smallest sum that covers this range. If the
		// smallest negative sum is larger than the current sum, then there is
		// no reason to include the current sum. This reasoning is often called
		// "Kadane's algorithm", but honestly, it's just basic logic.
		negSum := findPath(nums[l+1:j], k)
		sum += negSum
	}

	return res
}

// findPath finds the "cheapest" path through nums such that the stride is <= k
func findPath(nums []int, k int) int {
	// The optimal cost for a given index is given by not picking it and picking
	// the largest number in the prior k positions. This is a range-min-query
	// (RMQ), and I'm most comfortable implementing it as a segment tree.
	s, err := NewSegtree(math.MinInt64, max)
	if err != nil {
		panic(err)
	}
	n := len(nums)
	s.Reset(n + 2)
	s.Update(0, 0)
	for i := 1; i < n+2; i++ {
		lo := max(0, i-k)
		hi := i - 1
		cost := s.QueryRange(lo, hi)
		if i <= n {
			cost += nums[i-1]
		}
		s.Update(i, cost)
	}
	a := s.QueryRange(n+1, n+1)
	return a
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
