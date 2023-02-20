package p0218skylineproblemsegtree

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_getSkyline(t *testing.T) {
	for _, tc := range []struct {
		buildings [][]int
		want      [][]int
	}{
		{
			[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			leetcode.ParseMatrix("[[0,2,3],[2,5,3]]"),
			leetcode.ParseMatrix("[[0,3],[5,0]]"),
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.buildings), func(t *testing.T) {
			require.Equal(t, tc.want, getSkyline(tc.buildings))
		})
	}
}

func getSkyline(buildings [][]int) [][]int {
	// Steps to solving this problem:
	//
	// 1. Sort building edges by the x-coordinate
	// 2. Go through building edges, keeping track of which building has the
	// maximum y-coordinate for each x-coord.
	// 3. Decide for each position whether it is part of the skyline (its
	// x-coordinate is unique and the y-coordinate has changed since the last
	// edge).
	//
	// Sorting is easy. Keeping track of the highest building can be done in
	// many ways, the most common of which is to use a maxHeap. However, I will
	// instead use a segment tree.
	//
	// First sort by x-coord.
	type edge struct {
		i      int
		x      int
		height int
	}
	n := len(buildings)
	edges := make([]edge, 2*n)
	for i, b := range buildings {
		edges[i*2] = edge{i, b[0], b[2]}
		edges[i*2+1] = edge{i, b[1], 0}
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].x == edges[j].x {
			return edges[i].height < edges[j].height
		}
		return edges[i].x < edges[j].x
	})
	m := 1
	for m < n {
		m <<= 1
	}
	tree := make([]int, 2*m)
	var res [][]int

	// For each edge in x-order
	for _, e := range edges {
		// Update max-tree
		tree[n+e.i] = e.height
		for i := (n + e.i) / 2; i >= 1; i /= 2 {
			tree[i] = max(tree[i*2], tree[i*2+1])
		}
		h := tree[1]

		// Compress edges to only contain different x-coords and heights
		res = append(res, []int{e.x, h})
		if len(res) == 1 {
			continue
		}
		if res[len(res)-2][0] == res[len(res)-1][0] {
			res[len(res)-2][1] = h
			res = res[:len(res)-1]
		}
		for len(res) > 1 && res[len(res)-1][1] == res[len(res)-2][1] {
			res = res[:len(res)-1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
