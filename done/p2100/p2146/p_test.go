package p2146

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_highestRankedKItems(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]int
		pricing []int
		start   []int
		k       int
		want    [][]int
	}{
		{[][]int{{1, 0, 1}, {3, 5, 2}, {1, 0, 1}}, []int{2, 5}, []int{1, 1}, 9, [][]int{{1, 1}, {1, 2}, {1, 0}}},
		{[][]int{{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1}}, []int{2, 5}, []int{0, 0}, 3, [][]int{{0, 1}, {1, 1}, {2, 1}}},
		{[][]int{{1, 2, 0, 1}, {1, 3, 3, 1}, {0, 2, 5, 1}}, []int{2, 3}, []int{2, 3}, 2, [][]int{{2, 1}, {1, 2}}},
		{[][]int{{1, 1, 1}, {0, 0, 1}, {2, 3, 4}}, []int{2, 3}, []int{0, 0}, 3, [][]int{{2, 1}, {2, 0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, highestRankedKItems(tc.grid, tc.pricing, tc.start, tc.k))
		})
	}
}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])

	type pos struct {
		i, j int
	}

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	// BFS visits cur, adding positions to next
	cur := []pos{{start[0], start[1]}}
	seen[start[0]][start[1]] = true
	next := []pos{}

	type distAndPos struct {
		dist, i, j int
	}

	res := make([]distAndPos, 0, k)
	for dist := 0; len(cur) > 0 && len(res) < k; dist++ {
		next := next[:0] // re-use previous slice (note swap at end of for loop)

		// Visit each cell at the current distance
		for _, cell := range cur {
			if price := grid[cell.i][cell.j]; price >= pricing[0] && price <= pricing[1] {
				res = append(res, distAndPos{
					dist: dist,
					i:    cell.i,
					j:    cell.j,
				})
			}

			// Visit adjacent cells
			for _, d := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				ii := cell.i + d[0]
				jj := cell.j + d[1]
				if ii < 0 || ii >= m || jj < 0 || jj >= n ||
					grid[ii][jj] == 0 || seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				next = append(next, pos{ii, jj})
			}
		}
		cur, next = next, cur
	}

	// Sort result
	sort.Slice(res, func(i, j int) bool {
		p1, p2 := res[i], res[j]
		if p1.dist == p2.dist { // Distance
			if grid[p1.i][p1.j] == grid[p2.i][p2.j] { // Price
				if p1.i == p2.i { // Row
					return p1.j < p2.j // Column
				}
				return p1.i < p2.i
			}
			return grid[p1.i][p1.j] < grid[p2.i][p2.j]
		}
		return p1.dist < p2.dist
	})

	// Trim excess items if necessary
	if len(res) > k {
		res = res[:k]
	}

	// Convert to result type
	ret := make([][]int, len(res))
	for i := range res {
		ret[i] = []int{res[i].i, res[i].j}
	}

	return ret
}
