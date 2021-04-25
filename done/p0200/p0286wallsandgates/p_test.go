package p0286wallsandgates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wallsAndGates(t *testing.T) {
	for _, tc := range []struct {
		rooms [][]int
		want  [][]int
	}{
		{
			[][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}},
			[][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}},
		},
		{[][]int{{-1}}, [][]int{{-1}}},
		{[][]int{{2147483647}}, [][]int{{2147483647}}},
		{[][]int{{0}}, [][]int{{0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rooms), func(t *testing.T) {
			wallsAndGates(tc.rooms)
			require.Equal(t, tc.want, tc.rooms)
		})
	}
}

const (
	wall  = -1
	gate  = 0
	empty = 1<<31 - 1
)

type pos struct {
	i, j int
}

func wallsAndGates(rooms [][]int) {
	// Perform BFS from each gate and mark the distance
	gates := make([]pos, 0)
	m, n := len(rooms), len(rooms[0])
	seen := make([][]bool, m)
	for i, row := range rooms {
		seen[i] = make([]bool, n)
		for j, room := range row {
			if room == gate {
				gates = append(gates, pos{i, j})
				seen[i][j] = true
			} else if room == wall {
				seen[i][j] = true
			}
		}
	}

	levelItems := [2][]pos{}
	levelItems[0] = gates
	var curLevel int
	for len(levelItems[curLevel%2]) > 0 {
		nextLevel := (curLevel + 1)
		levelItems[nextLevel%2] = levelItems[nextLevel%2][:0] // zero

		for _, p := range levelItems[curLevel%2] {
			// Add all unseen empty rooms
			for _, near := range []pos{
				{p.i - 1, p.j}, {p.i + 1, p.j}, {p.i, p.j - 1}, {p.i, p.j + 1},
			} {
				if near.i < 0 || near.i >= m || near.j < 0 || near.j >= n || seen[near.i][near.j] {
					continue
				}
				rooms[near.i][near.j] = nextLevel
				levelItems[nextLevel%2] = append(levelItems[nextLevel%2], near)
				seen[near.i][near.j] = true
			}
		}
		curLevel = nextLevel
	}
}
