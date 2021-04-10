package p1632ranktransformofmatrix

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matrixRankTransform(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   [][]int
	}{
		// {[][]int{{7, 7}, {7, 7}}, [][]int{{1, 1}, {1, 1}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {2, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, matrixRankTransform(tc.matrix))
		})
	}
}

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	matrixEntries := make([]matrixEntry, 0, m*n)
	for i := range matrix {
		for j, val := range matrix[i] {
			matrixEntries = append(matrixEntries, matrixEntry{i, j, val})
		}
	}

	// Create result matrix
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	sort.Slice(matrixEntries, func(i, j int) bool {
		return matrixEntries[i].val < matrixEntries[j].val
	})

	// Keep track of the maximum rank in each col / row
	rowMaxRank := make([]int, m)
	colMaxRank := make([]int, n)

	// for each matrix entry (from low to high)
	for i := 0; i < m*n; {
		// gather all entries with the same value
		var j int
		for j = i + 1; j < m*n && matrixEntries[j].val == matrixEntries[i].val; j++ {
		}
		dsu := NewDSU(m + n)
		for _, e := range matrixEntries[i:j] {
			// join column and row together
			dsu.union(e.i, m+e.j)
		}
		// for each set in the disjoint set union, the rank is the
		// maximum value found in rowMaxRank / colMaxRank + 1
		idxToRoot := make([]int, j-i)
		ranks := make(map[int]int)
		for i, e := range matrixEntries[i:j] {
			r := dsu.find(e.i)
			idxToRoot[i] = r
			ranks[r] = max(ranks[r], 1+rowMaxRank[e.i])
			ranks[r] = max(ranks[r], 1+colMaxRank[e.j])
		}

		for idx, root := range idxToRoot {
			e := matrixEntries[i+idx]
			res[e.i][e.j] = ranks[root]
			rowMaxRank[e.i] = ranks[root]
			colMaxRank[e.j] = ranks[root]
		}
		i = j
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type matrixEntry struct {
	i, j, val int
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b int) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		if d.size[a] < d.size[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
	}
}
