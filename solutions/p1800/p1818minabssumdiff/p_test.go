package p1818minabssumdiff

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minAbsoluteSumDiff(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 7, 5}, []int{2, 3, 5}, 3},
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, 0},
		{[]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}, 20},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, minAbsoluteSumDiff(tc.nums1, tc.nums2))
		})
	}
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	type diff struct {
		n1, n2, d int
	}
	diffs := make([]diff, n)
	var totalDiff int
	for i, n1 := range nums1 {
		n2 := nums2[i]
		d := abs(n2 - n1)
		totalDiff += d
		diffs[i] = diff{n1, n2, d}
	}
	sort.Slice(diffs, func(i, j int) bool {
		return diffs[i].d > diffs[j].d
	})
	sort.Ints(nums1)

	// Find improvements to the diffs until the diff is smaller
	// than the maximum improvement so far
	minDiff := totalDiff
	for _, d := range diffs {
		if minDiff < totalDiff-d.d {
			break
		}
		// Find max reduction possible for this location
		i := sort.Search(n, func(i int) bool {
			return d.n2 < nums1[i]
		})
		switch {
		case i == n:
			minDiff = min(minDiff, totalDiff-d.d+abs(d.n2-nums1[n-1]))
		case i == 0:
			minDiff = min(minDiff, totalDiff-d.d+abs(d.n2-nums1[0]))
		default:
			for l := i; l >= i-1; l-- {
				minDiff = min(minDiff, totalDiff-d.d+abs(d.n2-nums1[l]))
			}
			for r := i + 1; r < n && r <= i+1; r++ {
				minDiff = min(minDiff, totalDiff-d.d+abs(d.n2-nums1[r]))
			}
		}
	}
	return minDiff % 1000000007
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
