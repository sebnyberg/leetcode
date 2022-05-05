package p0718maximumlengthofrepeatedsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLength(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, findLength(tc.nums1, tc.nums2))
		})
	}
}

func findLength(nums1 []int, nums2 []int) int {
	// This exercise is not a typical sub-array problem but rather a sequence
	// matching problem. Typically sequence matching (e.g. edit distance) is done
	// with either rolling hash or dynamic programming.
	//
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	var maxVal int
	for i, n := range nums1 {
		for j, m := range nums2 {
			if n == m {
				val := dp[i][j] + 1
				dp[i+1][j+1] = val
				if val > maxVal {
					maxVal = val
				}
			}
		}
	}
	return maxVal
}
