package p0329longestincreasingpathinmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestIncreasingPath(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{
			{9, 9, 4},
			{6, 6, 8},
			{2, 1, 1},
		}, 4},
		{[][]int{
			{1},
		}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, longestIncreasingPath(tc.matrix))
		})
	}
}

func longestIncreasingPath(matrix [][]int) int {
	// DFS
	var maxPath int
	m, n := len(matrix), len(matrix[0])
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[i] {
			maxPath = max(maxPath, dfs(mem, matrix, m, n, i, j))
		}
	}
	return maxPath
}

func dfs(mem [][]int, matrix [][]int, m, n int, i, j int) int {
	var maxVal int
	for _, near := range [][2]int{
		{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
	} {
		nearI, nearJ := near[0], near[1]
		if nearI < 0 || nearJ < 0 || nearI >= m || nearJ >= n {
			continue
		}
		if matrix[nearI][nearJ] > matrix[i][j] {
			if mem[nearI][nearJ] == 0 {
				mem[nearI][nearJ] = dfs(mem, matrix, m, n, nearI, nearJ)
			}
			maxVal = max(maxVal, mem[nearI][nearJ])
		}
	}
	return 1 + maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
