package p0565arraynesting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_arrayNesting(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
		{[]int{0, 1, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, arrayNesting(tc.nums))
		})
	}
}

func arrayNesting(nums []int) int {
	// Search for the longest chain using DFS and memoization
	mem := make(map[int]int, len(nums))
	visited := make([]bool, len(nums))
	var maxRes int
	for i := 0; i < len(nums); i++ {
		res := dfs(mem, visited, nums, i)
		maxRes = max(maxRes, res)
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(mem map[int]int, visited []bool, nums []int, idx int) int {
	if visited[idx] {
		return 0
	}
	if _, exists := mem[idx]; !exists {
		visited[idx] = true
		n := 1 + dfs(mem, visited, nums, nums[idx])
		visited[idx] = false
		mem[idx] = n
	}
	return mem[idx]
}
