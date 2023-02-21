package p0327countofrangesumsegtree

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countRangeSum(t *testing.T) {
	for _, tc := range []struct {
		nums         []int
		lower, upper int
		want         int
	}{
		{[]int{1, 2, -3, -3, -1, -3}, 2, 4, 2},
		{[]int{-2, 5, -1}, -2, 2, 3},
		{[]int{0}, 0, 0, 1},
		{[]int{5, -23, -5, -1, -21, 13, 15, 7, 18, 4, 7, 26, 29, -7, -28, 11, -20, -29, 19, 22, 15, 25, 17, -13, 11, -15, 19, -8, 3, 12, -1, 2, -1, -21, -10, -7, 14, -12, -14, -8, -1, -30, 19, -27, 16, 2, -15, 23, 6, 14, 23, 2, -4, 4, -9, -8, 10, 20, -29, 29},
			-19,
			10,
			362,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countRangeSum(tc.nums, tc.lower, tc.upper))
		})
	}
}

func countRangeSum(nums []int, lower int, upper int) int {
	// Compute a prefix sum of nums.
	m := len(nums)
	presum := make([]int, m+1)

	// Collect unique possible sums in a sorted list
	summ := make(map[int]struct{})
	summ[0] = struct{}{}
	for i := range nums {
		presum[i+1] = presum[i] + nums[i]
		summ[presum[i+1]] = struct{}{}
	}
	suml := make([]int, 0, len(summ))
	for sum := range summ {
		suml = append(suml, sum)
	}
	sort.Ints(suml)

	// Map each prefix sum to its index
	sumi := make(map[int]int) // map presum to sorted 1d index
	for i := range suml {
		sumi[suml[i]] = i
	}
	n := 1
	for bits.OnesCount(uint(len(suml))) > 1 {
		suml = append(suml, math.MaxInt64)
	}
	for n < len(suml) {
		n <<= 1
	}

	// Segment tree
	//
	// The tree will contain a count prefix sums. Since the sums are added as
	// elements are visited from left-to-right, the tree can be queried to
	// quickly count the number of prior sums that could be removed from the
	// current one to form a valid subarray.
	tree := make([]int, n*2)

	// update increases the count of the presum x in the tree
	update := func(tree []int, presum int) {
		i := sumi[presum]
		for j := i + n; j >= 1; j /= 2 {
			tree[j] += 1
		}
	}

	// query counts the number of prefix sums in the range [qlo, qhi]
	var query func(tree []int, i, loidx, hiidx, qlo, qhi int) int
	query = func(tree []int, i, loidx, hiidx, qlo, qhi int) int {
		lo := suml[loidx]
		hi := suml[hiidx]
		if qhi < lo || qlo > hi {
			return 0
		}
		if qlo <= lo && qhi >= hi {
			return tree[i]
		}
		mid := loidx + (hiidx-loidx)/2
		return query(tree, i*2, loidx, mid, qlo, qhi) +
			query(tree, i*2+1, mid+1, hiidx, qlo, qhi)
	}

	update(tree, 0) // zero-sum is always available

	var res int
	for i := range nums {
		// A prior sum between [lo, hi) can be added to the current sum to form
		// a valid subarray that ends in nums[i]
		sum := presum[i+1]
		lo := sum - upper
		hi := sum - lower
		res += query(tree, 1, 0, n-1, lo, hi)
		update(tree, sum)
	}
	return res
}
