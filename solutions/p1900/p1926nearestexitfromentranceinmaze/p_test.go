package p1926nearestexitfromentranceinmaze

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nearestExit(t *testing.T) {
	for i, tc := range []struct {
		maze     [][]byte
		entrance []int
		want     int
	}{
		{[][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		}, []int{1, 2}, 1},
		{[][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		}, []int{1, 0}, 2},
		{[][]byte{{'.', '+'}}, []int{0, 0}, -1},
	} {
		t.Run(fmt.Sprintf("%v, %+v", i, tc.maze), func(t *testing.T) {
			require.Equal(t, tc.want, nearestExit(tc.maze, tc.entrance))
		})
	}
}

func nearestExit(maze [][]byte, entrance []int) int {
	type position struct {
		i, j int
	}
	visited := make(map[position]bool)
	visited[position{entrance[0], entrance[1]}] = true
	tovisit := []position{{entrance[0], entrance[1]}}
	next := []position{}
	m, n := len(maze), len(maze[0])
	steps := 1
	for len(tovisit) > 0 {
		next = next[:0]
		for _, p := range tovisit {
			for _, near := range []position{
				{p.i + 1, p.j}, {p.i - 1, p.j}, {p.i, p.j - 1}, {p.i, p.j + 1},
			} {
				if visited[near] || near.i < 0 || near.j < 0 || near.i >= m || near.j >= n || maze[near.i][near.j] == '+' {
					continue
				}
				visited[near] = true
				if near.i == 0 || near.j == 0 || near.i == m-1 || near.j == n-1 {
					return steps
				}
				next = append(next, near)
			}
		}
		steps++
		next, tovisit = tovisit, next
	}
	return -1
}
