package p1035uncrossedlines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxUncrossedLines(t *testing.T) {
	for i, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			[]int{1, 4, 2},
			[]int{1, 2, 4},
			2,
		},
		{
			[]int{2, 5, 1, 2, 5},
			[]int{10, 5, 2, 1, 5, 2},
			3,
		},
		{
			[]int{1, 3, 7, 1, 7, 5},
			[]int{1, 9, 2, 5, 1},
			2,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxUncrossedLines(tc.nums1, tc.nums2))
		})
	}
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	mem := make(map[[2]int]int)
	res := dfs(mem, nums1, nums2, 0, 0, m, n)
	return res
}

func dfs(mem map[[2]int]int, nums1, nums2 []int, i, j, m, n int) int {
	if i >= m || j >= n {
		return 0
	}
	key := [2]int{i, j}
	if v, exists := mem[key]; exists {
		return v
	}
	if nums1[i] == nums2[j] {
		mem[key] = 1 + dfs(mem, nums1, nums2, i+1, j+1, m, n)
		return mem[key]
	}
	// Try to match nothing on both sides
	res := dfs(mem, nums1, nums2, i+1, j, m, n)
	res = max(res, dfs(mem, nums1, nums2, i, j+1, m, n))

	// Or match the current to any other one
	for k := j + 1; k < len(nums2); k++ {
		if nums1[i] == nums2[k] {
			res = max(res, 1+dfs(mem, nums1, nums2, i+1, k+1, m, n))
		}
	}
	for k := i + 1; k < len(nums1); k++ {
		if nums2[j] == nums1[k] {
			res = max(res, 1+dfs(mem, nums1, nums2, k+1, j+1, m, n))
		}
	}
	mem[key] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
