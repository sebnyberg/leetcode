package p2333minimumsumofsquareddifference

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSumSquareDiff(t *testing.T) {
	for _, tc := range []struct {
		nums1  []int
		nums2  []int
		k1, k2 int
		want   int64
	}{
		{[]int{7, 5, 0, 12, 14}, []int{7, 5, 0, 12, 14}, 2, 9, 0},
		{[]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1, 43},
		{[]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0, 579},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, minSumSquareDiff(tc.nums1, tc.nums2, tc.k1, tc.k2))
		})
	}
}

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	// There is no point to negative integers, because it would just
	// result in the same value anyway
	// The goal is simply to reduce the largest difference each time.
	n := len(nums1)
	deltas := make([]int, n, n+1)
	var sum int
	for i := range nums1 {
		deltas[i] = abs(nums1[i] - nums2[i])
		sum += deltas[i]
	}
	k := k1 + k2
	if k >= sum {
		return 0
	}

	deltas = append(deltas, 0) // sentinel
	sort.Slice(deltas, func(i, j int) bool {
		return deltas[i] > deltas[j]
	})
	// Each delta can be reduced either with k1 or k2
	// Reduce until we run out of k, or all deltas are zero.
	for i := 1; k > 0; i++ {
		// Attempt to reduce all prior numbers to match this one
		shouldReduce := i * (deltas[i-1] - deltas[i])

		// Possible! Skip this number
		if k >= shouldReduce {
			k -= shouldReduce
			continue
		}

		// Not possible. Reduce prior numbers as much as possible
		m := k / i
		rest := k % i
		actual := deltas[i-1] - m
		for j := i - 1; j >= 0; j-- {
			deltas[j] = actual
			if rest > 0 {
				deltas[j] -= 1
				rest--
			}
		}
		k = 0
	}
	// Calculate total sum of squares
	var res int64
	for i := 0; i < n; i++ {
		res += int64(deltas[i] * deltas[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
