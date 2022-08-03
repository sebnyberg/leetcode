package p2285maximumtotalimportanceofroads

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumImportance(t *testing.T) {
	for _, tc := range []struct {
		n     int
		roads [][]int
		want  int64
	}{
		{5, leetcode.ParseMatrix("[[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]]"), 43},
		{5, leetcode.ParseMatrix("[[0,3],[2,4],[1,3]]"), 20},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maximumImportance(tc.n, tc.roads))
		})
	}
}

func maximumImportance(n int, roads [][]int) int64 {
	// The goal is to rank each city by the number of edges, then assign values
	// accordingly.
	// Once values are assigned, visit each edge (given by roads) and count the
	// result.
	type edge struct {
		idx   int
		count int
	}
	edgeCount := make([]edge, n)
	for i := range edgeCount {
		edgeCount[i].idx = i
	}
	for _, r := range roads {
		a, b := r[0], r[1]
		edgeCount[a].count++
		edgeCount[b].count++
	}
	sort.Slice(edgeCount, func(i, j int) bool {
		return edgeCount[i].count < edgeCount[j].count
	})
	cityVal := make([]int, n)
	for i, e := range edgeCount {
		cityVal[e.idx] = i + 1
	}
	var res int64
	for _, r := range roads {
		res += int64(cityVal[r[0]] + cityVal[r[1]])
	}
	return res
}
