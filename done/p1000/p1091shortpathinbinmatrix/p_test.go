package p1091shortpathinbinmatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{[][]int{{0, 1}, {1, 0}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			got := shortestPathBinaryMatrix(tc.grid)
			require.Equal(t, tc.want, got)
		})
	}
}

type pos struct{ i, j int }

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m := len(grid)
	n := len(grid[0])
	available := make([][]bool, m+2)
	for i := range available {
		available[i] = make([]bool, n+2)
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				available[i+1][j+1] = true
			}
		}
	}

	todo := map[pos]struct{}{
		{1, 1}: {},
	}

	steps := 1
	for len(todo) > 0 {
		// for each position in todo, pop it and add all nearby
		// available positions to the list of todos
		for todoPos := range todo {
			if todoPos.i == m && todoPos.j == n {
				return steps
			}
			available[todoPos.i][todoPos.j] = false
		}
		newTodo := make(map[pos]struct{})
		for t := range todo {
			for _, nearby := range []pos{
				{t.i - 1, t.j - 1}, {t.i - 1, t.j}, {t.i - 1, t.j + 1},
				{t.i, t.j - 1}, {t.i, t.j}, {t.i, t.j + 1},
				{t.i + 1, t.j - 1}, {t.i + 1, t.j}, {t.i + 1, t.j + 1},
			} {
				if available[nearby.i][nearby.j] {
					newTodo[nearby] = struct{}{}
				}
			}
		}
		todo = newTodo
		steps++
	}

	return -1
}
