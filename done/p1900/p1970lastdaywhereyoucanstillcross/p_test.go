package p1970lastdaywhereyoucanstillcross

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_latestDayToCross(t *testing.T) {
	for _, tc := range []struct {
		row   int
		col   int
		cells [][]int
		want  int
	}{
		{6, 2, [][]int{{4, 2}, {6, 2}, {2, 1}, {4, 1}, {6, 1}, {3, 1}, {2, 2}, {3, 2}, {1, 1}, {5, 1}, {5, 2}, {1, 2}}, 3},
		{2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, 2},
		{2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 1},
		{3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.row), func(t *testing.T) {
			require.Equal(t, tc.want, latestDayToCross(tc.row, tc.col, tc.cells))
		})
	}
}

func latestDayToCross(row int, col int, cells [][]int) int {
	// Instead of removing land, add land by going in reverse.
	// Add an auxiliary set of cells above and below the top and bottom rows.
	// Use DSU to determine if the top part is connected to the bottom part.
	to1d := func(i, j int) int {
		return i*col + j
	}

	parent := make([]int, (row+2)*(col))
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			root := find(parent[a])
			parent[a] = root
		}
		return parent[a]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			if rb < ra {
				ra, rb = rb, ra
			}
			parent[rb] = ra
		}
	}

	// Setup a strip of land at the top and bottom
	land := make([][]bool, row+2)
	for i := range land {
		land[i] = make([]bool, col)
	}
	for j := 0; j < col; j++ {
		land[0][j] = true
		land[row+1][j] = true
		if j > 0 {
			union(to1d(0, j), 0)
			union(to1d(row+1, j), to1d(row+1, j-1))
		}
	}

	near := func(i, j int) [][2]int {
		return [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < row+2 && j >= 0 && j < col && land[i][j]
	}

	// Add land until find(some bottom cell) == 0
	for day := len(cells) - 1; day >= 0; day-- {
		i, j := cells[day][0], cells[day][1]-1
		land[i][j] = true
		for _, nei := range near(i, j) {
			if ok(nei[0], nei[1]) {
				union(to1d(i, j), to1d(nei[0], nei[1]))
			}
		}
		if find(to1d(row+1, col-1)) == 0 {
			return day
		}
	}

	return 0
}
