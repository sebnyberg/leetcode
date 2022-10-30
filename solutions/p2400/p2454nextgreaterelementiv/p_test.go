package p2454nextgreaterelementiv

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_secondGreaterElement(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{1, 17, 18, 0, 18, 10, 20, 0}, []int{18, 18, -1, 10, -1, -1, -1, -1}},
		{[]int{2, 4, 0, 9, 6}, []int{9, 6, 6, -1, -1}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, secondGreaterElement(tc.nums))
		})
	}
}

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		ra := find(parent[a])
		parent[a] = ra
		return ra
	}
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		if nums[indices[i]] == nums[indices[j]] {
			return indices[i] > indices[j]
		}
		return nums[indices[i]] < nums[indices[j]]
	})
	res := make([]int, n)
	for _, idx := range indices {
		if idx >= n-2 {
			res[idx] = -1
			nums[idx] = -1
			parent[idx] = n - 1
			continue
		}
		next := find(idx + 1)
		parent[idx] = next
		if next == n-1 {
			res[idx] = -1
			continue
		}
		res[idx] = nums[find(next+1)]
	}
	return res
}
