package p2215findthedifferenceoftwoarrays

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDifference(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{
			[]int{1, 2, 3},
			[]int{2, 4, 6},
			leetcode.ParseMatrix("[[1,3],[4,6]]"),
		},
		{
			[]int{1, 2, 3, 3},
			[]int{1, 1, 2, 2},
			leetcode.ParseMatrix("[[3],[]]"),
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, findDifference(tc.nums1, tc.nums2))
		})
	}
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := make(map[int]struct{}), make(map[int]struct{})
	for _, n1 := range nums1 {
		m1[n1] = struct{}{}
	}

	for _, n2 := range nums2 {
		m2[n2] = struct{}{}
	}
	var res1 []int
	for n1 := range m1 {
		if _, exists := m2[n1]; !exists {
			res1 = append(res1, n1)
		}
	}
	var res2 []int
	for n2 := range m2 {
		if _, exists := m1[n2]; !exists {
			res2 = append(res2, n2)
		}
	}
	return [][]int{res1, res2}
}
