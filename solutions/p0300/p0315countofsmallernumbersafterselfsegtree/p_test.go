package p0315countofsmallernumbersafterself

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSmaller(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1}, []int{0}},
		{[]int{-1, -1}, []int{0, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countSmaller(tc.nums))
		})
	}
}

func countSmaller(nums []int) []int {
	// We can use a segment-tree to store counts per value
	minval := math.MaxInt32
	for _, x := range nums {
		minval = min(minval, x)
	}
	maxval := math.MinInt32
	for i := range nums {
		nums[i] += -minval
		maxval = max(maxval, nums[i])
	}
	n := 1
	for n < maxval+1 {
		n <<= 1
	}
	tree := make([]int, 2*n)
	// Query returns the count of numbers in the interval [qlo, qhi)
	var query func(i, lo, hi, qlo, qhi int) int
	query = func(i, lo, hi, qlo, qhi int) int {
		if hi <= qlo || lo >= qhi {
			return 0
		}
		if qlo <= lo && qhi >= hi {
			return tree[i]
		}
		mid := lo + (hi-lo)/2
		return query(i*2, lo, mid, qlo, qhi) +
			query(i*2+1, mid, hi, qlo, qhi)
	}
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		// Add number to segtree
		for j := nums[i] + n; j >= 1; j /= 2 {
			tree[j] += 1
		}

		// Query current value
		res[i] = query(1, 0, n, 0, nums[i])
	}
	return res
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
