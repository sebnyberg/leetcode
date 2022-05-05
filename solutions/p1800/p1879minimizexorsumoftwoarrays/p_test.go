package p1879minimizexorsumoftwoarrays

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minXORSum(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2}, []int{2, 3}, 2},
		{[]int{1, 0, 3}, []int{5, 3, 4}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, minimumXORSum(tc.nums1, tc.nums2))
		})
	}
}

func minimumXORSum(nums1 []int, nums2 []int) int {
	// All permutations of nums1 is up to 14!
	// Possibly OK to brute-force
	n := len(nums1)
	mem := make([]int, 1<<14)
	for i := range mem {
		mem[i] = math.MaxInt32
	}
	return perm(mem, nums1, nums2, 0, n, 0)
}

func perm(mem []int, nums1, nums2 []int, i, n, mask int) int {
	if i == n {
		return 0
	}
	if mem[mask] == math.MaxInt32 {
		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				mem[mask] = min(mem[mask],
					nums1[i]^nums2[j]+perm(mem, nums1, nums2, i+1, n, mask+(1<<j)),
				)
			}
		}
		return mem[mask]
	}
	return mem[mask]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
