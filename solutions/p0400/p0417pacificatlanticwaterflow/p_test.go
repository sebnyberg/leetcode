package p0417pacificatlanticwaterflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pacificAtlantic(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{10, 10, 10}, {10, 1, 10}, {10, 10, 10}},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		},
		// {
		// 	[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
		// 	[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		// },
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Subset(t, tc.want, pacificAtlantic(tc.matrix))
		})
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	n, m := len(matrix[0]), len(matrix)
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	type position struct {
		i, j int
	}

	// Pacific
	curr := make([]position, 0, n+m-1)
	for i := range pacific {
		pacific[i][0] = true
		curr = append(curr, position{i, 0})
	}
	for i := range pacific[0] {
		pacific[0][i] = true
		curr = append(curr, position{0, i})
	}

	next := []position{}
	for len(curr) > 0 {
		next = next[:0]
		for _, t := range curr {
			for _, pos := range []position{
				{t.i, t.j - 1},
				{t.i, t.j + 1},
				{t.i - 1, t.j},
				{t.i + 1, t.j},
			} {
				if pos.i < 0 || pos.i >= m || pos.j < 0 || pos.j >= n {
					continue
				}
				if !pacific[pos.i][pos.j] && matrix[pos.i][pos.j] >= matrix[t.i][t.j] {
					next = append(next, pos)
					pacific[pos.i][pos.j] = true
				}
			}
		}
		curr, next = next, curr
	}

	// Atlantic
	for i := range atlantic {
		atlantic[i][n-1] = true
		curr = append(curr, position{i, n - 1})
	}
	for i := range atlantic[0] {
		atlantic[m-1][i] = true
		curr = append(curr, position{m - 1, i})
	}

	for len(curr) > 0 {
		newTodos := make([]position, 0)
		for _, t := range curr {
			for _, pos := range []position{
				{t.i - 1, t.j},
				{t.i + 1, t.j},
				{t.i, t.j - 1},
				{t.i, t.j + 1},
			} {
				if pos.i < 0 || pos.i >= m || pos.j < 0 || pos.j >= n {
					continue
				}
				if !atlantic[pos.i][pos.j] && matrix[pos.i][pos.j] >= matrix[t.i][t.j] {
					newTodos = append(newTodos, pos)
					atlantic[pos.i][pos.j] = true
				}
			}
		}
		curr = newTodos
	}

	res := make([][]int, 0)
	for i := range pacific {
		for j, reach := range pacific[i] {
			if reach && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
