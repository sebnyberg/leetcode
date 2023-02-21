package p1521findavalueofamysteriousfunctionclosesttotarget

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closestToTarget(t *testing.T) {
	for i, tc := range []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{9, 12, 3, 7, 15}, 5, 2},
		{[]int{1000000, 1000000, 1000000}, 1, 999999},
		{[]int{1, 2, 4, 8, 16}, 0, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, closestToTarget(tc.arr, tc.target))
		})
	}
}

func closestToTarget(arr []int, target int) int {
	// The optimal subarray must start with some index. If the best possible
	// ending position can be found in ~O(log*n), then the total time complexity
	// is good enough (less than O(n^2)).
	//
	// Usually, you can use a prefix-sum to calculate such segments, but it will
	// not work for bitwise-or. Instead, a segment tree can be used.
	//
	// Note that there are probably more efficient ways to do this - I'm using a
	// segment tree because I'm practicing it.
	n := 1
	for n < len(arr) {
		n *= 2
	}

	// Build the tree
	tree := make([]int, n*2)
	for i := range arr {
		tree[n+i] = arr[i]
	}
	for i := n + len(arr); i < len(tree); i++ {
		tree[i] = math.MaxInt32
	}
	for i := n - 1; i >= 1; i-- {
		tree[i] = tree[2*i] & tree[2*i+1]
	}

	var q func(i, lo, hi, qlo, qhi int) int
	q = func(i, lo, hi, qlo, qhi int) int {
		if qlo <= lo && qhi >= hi {
			return tree[i]
		}
		if qhi < lo || qlo > hi {
			return math.MaxInt64 // all 1's
		}
		mid := lo + (hi-lo)/2
		return q(i*2, lo, mid, qlo, qhi) &
			q(i*2+1, mid+1, hi, qlo, qhi)
	}

	res := math.MaxInt32
	for i := range arr {
		// Find the shortest subarray that is smaller than target
		lo, hi := i, len(arr)-1
		for lo != hi {
			mid := lo + (hi-lo)/2
			midval := q(1, 0, n-1, i, mid)
			if midval >= target {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		above := q(1, 0, n-1, i, lo)
		res = min(res, abs(target-above))
		if lo > i {
			below := q(1, 0, n-1, i, lo-1)
			res = min(res, abs(target-below))
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
