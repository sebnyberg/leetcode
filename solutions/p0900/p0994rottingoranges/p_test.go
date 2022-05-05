package p0994rottingoranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_orangesRotting(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, orangesRotting(tc.grid))
		})
	}
}

func orangesRotting(grid [][]int) int {
	type pos struct{ i, j int }
	m, n := len(grid), len(grid[0])
	cur := make([]pos, 0, m)
	var freshCount int
	var seen [10][10]bool
	for i := range grid {
		for j, val := range grid[i] {
			switch val {
			case 1:
				freshCount++
			case 2:
				cur = append(cur, pos{i, j})
				seen[i][j] = true
			}
		}
	}
	validPos := func(p pos) bool {
		return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n
	}
	next := make([]pos, 0, m)
	var minutes int
	for freshCount > 0 && len(cur) > 0 {
		next = next[:0]
		for _, p := range cur {
			for _, near := range []pos{
				{p.i + 1, p.j}, {p.i - 1, p.j}, {p.i, p.j - 1}, {p.i, p.j + 1},
			} {
				if !validPos(near) || seen[near.i][near.j] || grid[near.i][near.j] != 1 {
					continue
				}
				seen[near.i][near.j] = true
				freshCount--
				next = append(next, near)
			}
		}
		minutes++
		cur, next = next, cur
	}
	if freshCount != 0 {
		return -1
	}
	return minutes
}
