package p2392buildamatrixwithconditions

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_buildMatrix(t *testing.T) {
	for _, tc := range []struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
		want          [][]int
	}{
		{
			3,
			leetcode.ParseMatrix("[[1,2],[3,2]]"),
			leetcode.ParseMatrix("[[2,1],[3,2]]"),
			leetcode.ParseMatrix("[[3,0,0],[0,0,1],[0,2,0]]"),
		},
		{
			3,
			leetcode.ParseMatrix("[[1,2],[2,3],[3,1],[2,3]]"),
			leetcode.ParseMatrix("[[2,1]]"),
			[][]int{},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, buildMatrix(tc.k, tc.rowConditions, tc.colConditions))
		})
	}
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	adj := make([][]int, k+1)
	indeg := make([]int, k+1)
	curr := []int{}
	next := []int{}

	// We can topo-sort rows / columns to get the answers
	topoSort := func(conditions [][]int) []int {
		for i := range indeg {
			indeg[i] = 0
			adj[i] = adj[i][:0]
		}
		for _, c := range conditions {
			above, below := c[0], c[1]
			if above == below {
				return nil
			}
			adj[above] = append(adj[above], below)
			indeg[below]++
		}
		curr = curr[:0]
		next = next[:0]
		res := make([]int, 0, k)
		for i := 1; i < len(indeg); i++ {
			if indeg[i] == 0 {
				curr = append(curr, i)
				res = append(res, i)
			}
		}
		for len(curr) > 0 {
			next = next[:0]
			for _, a := range curr {
				for _, b := range adj[a] {
					indeg[b]--
					if indeg[b] == 0 {
						res = append(res, b)
						next = append(next, b)
					}
				}
			}
			curr, next = next, curr
		}
		return res
	}
	rows := topoSort(rowConditions)
	if len(rows) != k {
		return [][]int{}
	}
	cols := topoSort(colConditions)
	if len(cols) != k {
		return [][]int{}
	}
	colIdx := make([]int, k+1)
	for i, c := range cols {
		colIdx[c] = i
	}
	res := make([][]int, k)
	for i := range res {
		res[i] = make([]int, k)
		res[i][colIdx[rows[i]]] = rows[i]
	}
	return res
}
