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

	var seen [10001][10001]bool

	canVisit := func(p pos) bool {
		return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n &&
			grid[p.i][p.j] != 0 && !seen[p.i][p.j]
	}
	wantsToBuy := func(p pos) bool {
		price := grid[p.i][p.j]
		return price >= pricing[0] && price <= pricing[1]
	}

	// BFS visits cur, adding positions to next
	cur := []pos{{start[0], start[1]}}
	next := []pos{}

	res := make([][]int, 0, k)
	for len(cur) > 0 {
		next := next[:0] // re-use previous slice (note swap at end of for loop)

		// Cur is sorted by price, row, col
		for _, p := range cur {
			if wantsToBuy(p) {
				res = append(res, []int{p.i, p.j})
				k--
				if k == 0 {
					return res
				}
			}

			// Visit adjacent neighbours
			for _, nei := range []pos{
				{p.i + 1, p.j}, {p.i - 1, p.j}, {p.i, p.j + 1}, {p.i, p.j - 1},
			} {
				if !canVisit(nei) {
					continue
				}
				seen[nei.i][nei.j] = true
				next = append(next, nei)
			}
		}

		// Sort alternatives
		sort.Slice(next, func(i, j int) bool {
			p1, p2 := next[i], next[j]
			if grid[p1.i][p1.j] == grid[p2.i][p2.j] { // Price
				if p1.i == p2.i { // Row
					return p1.j < p2.j // Column
				}
				return p1.i < p2.i
			}
			return grid[p1.i][p1.j] < grid[p2.i][p2.j]
		})

		cur, next = next, cur
	}
	return res
}
