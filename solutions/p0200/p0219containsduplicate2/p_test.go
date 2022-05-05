package p0219containsduplicate2

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1, 1, 2, 3}, 0, false},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, containsNearbyDuplicate(tc.nums, tc.k))
		})
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 1 {
		return false
	}
	type numIdx struct {
		n   int
		idx int
	}

	n := len(nums)
	numIndices := make([]numIdx, n)
	for i := range numIndices {
		numIndices[i] = numIdx{nums[i], i}
	}

	sort.Slice(numIndices, func(i, j int) bool { return numIndices[i].n < numIndices[j].n })

	var start int
	for start < n-1 {
		cur := numIndices[start].n
		for i := start; i < n && numIndices[i].n == cur; i++ {
			for j := 1; i+j < n && numIndices[i+j].n == cur; j++ {
				if abs(numIndices[i].idx-numIndices[i+j].idx) <= k {
					return true
				}
			}
		}
		for start < n && numIndices[start].n == cur {
			start++
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
