package p0870advantageshuffle

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_advantageCount(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		B    []int
		want []int
	}{
		{[]int{12, 24, 8, 32}, []int{13, 25, 32, 11}, []int{24, 32, 8, 12}},
		{[]int{2, 7, 11, 15}, []int{1, 10, 4, 11}, []int{2, 11, 7, 15}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.A, tc.B), func(t *testing.T) {
			require.Equal(t, tc.want, advantageCount(tc.A, tc.B))
		})
	}
}

type numSorter struct {
	nums    []int
	indices []int
}

func (s *numSorter) Swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	s.indices[i], s.indices[j] = s.indices[j], s.indices[i]
}

func (s numSorter) Len() int {
	return len(s.nums)
}

func (s numSorter) Less(i, j int) bool {
	return s.nums[i] < s.nums[j]
}

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	s := &numSorter{
		nums:    nums2,
		indices: make([]int, n),
	}
	for i := range nums2 {
		s.indices[i] = i
	}
	sort.Ints(nums1)
	sort.Sort(s)
	res := make([]int, n)
	any := make([]int, 0, n/10)
	var i, j int
	for i < n && j < n {
		if nums1[i] <= s.nums[j] {
			any = append(any, nums1[i])
			i++
		} else {
			res[s.indices[j]] = nums1[i]
			j++
			i++
		}
	}
	for k := j; k < n; k++ {
		res[s.indices[k]] = any[k-j]
	}
	return res
}
