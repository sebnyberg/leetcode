package p0934shortestbridge

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_shortestBridge(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,1,0],[0,0,0],[0,0,1]]"),
			2,
		},
		{
			leetcode.ParseMatrix("[[0,1],[1,0]]"),
			1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, shortestBridge(tc.grid))
		})
	}
}

func shortestBridge(grid [][]int) int {
	// first, mark one island with 2s
	curr := [][]int{}
	next := [][]int{}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	ok := func(i, j int) bool {
		return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
	}
	f := func() {
		for i := range grid {
			for j, v := range grid[i] {
				if v != 1 {
					continue
				}
				curr = curr[:0]
				curr = append(curr, []int{i, j})
				grid[i][j] = 2
				for len(curr) > 0 {
					next = next[:0]
					for _, x := range curr {
						for _, d := range dirs {
							ii := x[0] + d[0]
							jj := x[1] + d[1]
							if !ok(ii, jj) || grid[ii][jj] == 2 || grid[ii][jj] == 0 {
								continue
							}
							grid[ii][jj] = 2
							next = append(next, []int{ii, jj})
						}
					}
					curr, next = next, curr
				}
				return
			}
		}
	}
	f()

	// starting from a 1, expand until hitting a 2
	curr = curr[:0]
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 {
				curr = append(curr, []int{i, j})
			}
		}
	}
	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				if !ok(ii, jj) || grid[ii][jj] == 1 {
					continue
				}
				if grid[ii][jj] == 2 {
					return steps
				}
				grid[ii][jj] = 1
				next = append(next, []int{ii, jj})
			}
		}
		curr, next = next, curr
	}
	panic("hehe")
}
