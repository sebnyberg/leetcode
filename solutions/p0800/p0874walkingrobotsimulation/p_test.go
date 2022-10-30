package p0874walkingrobotsimulation

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_robotSim(t *testing.T) {
	for i, tc := range []struct {
		commands  []int
		obstacles [][]int
		want      int
	}{
		{[]int{2, 2, 5, -1, -1}, leetcode.ParseMatrix("[[-3,5],[-2,5],[3,2],[5,0],[-2,0],[-1,5],[5,-3],[0,0],[-4,4],[-3,4]]"), 81},
		{[]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}, 65},
		{[]int{4, -1, 3}, [][]int{}, 25},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, robotSim(tc.commands, tc.obstacles))
		})
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[[2]int]struct{})
	for _, o := range obstacles {
		m[[2]int{o[0], o[1]}] = struct{}{}
	}
	dirs := [][2]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	k := 0
	var y, x int
	var res int
	for _, cmd := range commands {
		switch cmd {
		case -2:
			k = (k + 3) % 4
		case -1:
			k = (k + 1) % 4
		default:
			for i := 0; i < cmd; i++ {
				nextY := y + dirs[k][1]
				nextX := x + dirs[k][0]
				if _, exists := m[[2]int{nextX, nextY}]; exists {
					break
				}
				y = nextY
				x = nextX
				res = max(res, x*x+y*y)
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
