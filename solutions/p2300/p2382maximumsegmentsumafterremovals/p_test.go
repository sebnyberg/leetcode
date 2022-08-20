package p2382maximumsegmentsumafterremovals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSegmentSums(t *testing.T) {
	for _, tc := range []struct {
		nums          []int
		removeQueries []int
		want          []int64
	}{
		{[]int{1, 2, 5, 6, 1}, []int{0, 3, 2, 4, 1}, []int64{14, 7, 2, 2, 0}},
		{[]int{3, 2, 11, 1}, []int{3, 2, 1, 0}, []int64{16, 5, 3, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSegmentSum(tc.nums, tc.removeQueries))
		})
	}
}

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {

	n := len(nums)
	parent := make([]int, n)
	segmentSums := make([]int, n)
	for i := range parent {
		parent[i] = i
		segmentSums[i] = nums[i]
	}

	blocked := make([]bool, len(nums))
	for _, q := range removeQueries {
		blocked[q] = true
	}

	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		root := find(parent[a])
		parent[a] = root
		return root
	}
	union := func(a, b int) int {
		var res int
		ra := find(a)
		rb := find(b)
		res = max(segmentSums[ra], segmentSums[rb])
		if ra != rb {
			parent[ra] = rb
			segmentSums[rb] += segmentSums[ra]
			res = max(res, segmentSums[rb])
		}
		return res
	}
	// Join all segments which are not blocked
	var maxSum int
	for i := range nums {
		if blocked[i] {
			continue
		}
		if i > 0 && !blocked[i-1] {
			union(i-1, i)
			maxSum = max(maxSum, segmentSums[find(i)])
		}
	}

	// Re-introduce values one by one
	res := make([]int64, len(removeQueries))
	for i := len(removeQueries) - 1; i >= 0; i-- {
		res[i] = int64(maxSum)
		q := removeQueries[i]
		blocked[q] = false
		maxSum = max(maxSum, nums[q])
		if q > 0 && !blocked[q-1] {
			maxSum = max(maxSum, union(q-1, q))
		}
		if q < n-1 && !blocked[q+1] {
			maxSum = max(maxSum, union(q, q+1))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
