package p0542z01matrix

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_updateMatrix(t *testing.T) {
	for _, tc := range []struct {
		mat  [][]int
		want [][]int
	}{
		{
			leetcode.ParseMatrix("[[0,0,0],[0,1,0],[0,0,0]]"),
			leetcode.ParseMatrix("[[0,0,0],[0,1,0],[0,0,0]]"),
		},
		{
			leetcode.ParseMatrix("[[0,0,0],[0,1,0],[1,1,1]]"),
			leetcode.ParseMatrix("[[0,0,0],[0,1,0],[1,2,1]]"),
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, updateMatrix(tc.mat))
		})
	}
}

func updateMatrix(mat [][]int) [][]int {
	// Simple BFS
	// Visit all zeroes in mat, marking non-zero neighbours as the next to visit.
	m, n := len(mat), len(mat[0])
	seen := make([][]bool, m)
	curr := make([][2]int, 0, m)
	dirs := [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && !seen[i][j]
	}
	for i := range mat {
		seen[i] = make([]bool, n)
		for j, v := range mat[i] {
			if v == 1 {
				continue
			}
			seen[i][j] = true
			curr = append(curr, [2]int{i, j})
		}
	}

	// Perform bfs
	next := make([][2]int, 0, m)
	for val := 1; len(curr) > 0; val++ {
		next = next[:0]

		for _, p := range curr {
			i, j := p[0], p[1]
			for _, d := range dirs {
				ii, jj := i+d[0], j+d[1]
				if !ok(ii, jj) {
					continue
				}
				next = append(next, [2]int{ii, jj})
				seen[ii][jj] = true
				mat[ii][jj] = val
			}
		}
		curr, next = next, curr
	}

	return mat
}
