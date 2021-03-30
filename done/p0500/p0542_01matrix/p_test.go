package p0542_01matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_updateMatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1}},
			[][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 2, 1, 1, 0, 1}, {2, 1, 1, 1, 1, 2, 1, 0, 1, 0}, {3, 2, 2, 1, 0, 1, 0, 0, 1, 1}},
		},
		// {[][]int{
		// 	{0, 0, 0},
		// 	{0, 1, 0},
		// 	{1, 1, 1},
		// }, [][]int{
		// 	{0, 0, 0},
		// 	{0, 1, 0},
		// 	{1, 2, 1},
		// }},
		// {[][]int{
		// 	{0, 0, 0},
		// 	{0, 1, 0},
		// 	{0, 0, 0},
		// }, [][]int{
		// 	{0, 0, 0},
		// 	{0, 1, 0},
		// 	{0, 0, 0},
		// }},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, updateMatrix(tc.matrix))
		})
	}
}

func updateMatrix(matrix [][]int) [][]int {
	visited := make([][]bool, len(matrix))
	res := make([][]int, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
		res[i] = make([]int, len(matrix[0]))
	}

	// Find zeroes in the matrix
	toVisit := make([]pos, 0)
	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				toVisit = append(toVisit, pos{i, j})
				visited[i][j] = true
			}
		}
	}

	var dist int
	for len(toVisit) > 0 {
		new := make([]pos, 0)
		for _, p := range toVisit {
			for _, nearby := range []pos{
				{p.i - 1, p.j}, {p.i + 1, p.j}, {p.i, p.j + 1}, {p.i, p.j - 1},
			} {
				if nearby.i < 0 || nearby.i >= len(matrix) || nearby.j < 0 || nearby.j >= len(matrix[0]) {
					continue
				}
				if !visited[nearby.i][nearby.j] {
					new = append(new, nearby)
					visited[nearby.i][nearby.j] = true
				}
			}
			res[p.i][p.j] = dist
		}
		dist++
		toVisit = new
	}

	return res
}

type pos struct {
	i, j int
}
