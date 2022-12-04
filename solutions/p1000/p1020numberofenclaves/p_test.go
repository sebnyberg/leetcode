package p1020numberofenclaves

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numEnclaves(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,0,0,1,1,1,0,1,0,0],[1,1,0,0,0,1,0,1,1,1],[0,0,0,1,1,1,0,1,0,0],[0,1,1,0,0,0,1,0,1,0],[0,1,1,1,1,1,0,0,1,0],[0,0,1,0,1,1,1,1,0,1],[0,1,1,0,0,0,1,1,1,1],[0,0,1,0,0,1,0,1,0,1],[1,0,1,0,1,1,0,0,0,0],[0,0,0,0,1,1,0,0,0,1]]"),
			3,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numEnclaves(tc.grid))
		})
	}
}

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	markZeroes := func(i, j int) int {
		curr := [][2]int{[2]int{i, j}}
		next := [][2]int{}
		var foundEdge bool
		grid[i][j] = 0
		res := 1
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				i := x[0]
				j := x[1]
				if i == 0 || j == 0 || i == m-1 || j == n-1 {
					foundEdge = true
				}
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || grid[ii][jj] == 0 {
						continue
					}
					grid[ii][jj] = 0
					res++
					next = append(next, [2]int{ii, jj})
				}
			}
			curr, next = next, curr
		}
		if foundEdge {
			return 0
		}
		return res
	}
	var res int
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 {
				res += markZeroes(i, j)
			}
		}
	}
	return res
}
