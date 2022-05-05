package p0797allpathsfromsourcetotaget

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_allPathsSourceTarget(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		want  [][]int
	}{
		{[][]int{{1, 2}, {3}, {3}, {}}, [][]int{{0, 1, 3}, {0, 2, 3}}},
		{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}},
		{[][]int{{1}, {}}, [][]int{{0, 1}}},
		{[][]int{{1, 2, 3}, {2}, {3}, {}}, [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			require.Equal(t, tc.want, allPathsSourceTarget(tc.graph))
		})
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	// Since the exercise is not counting paths, the solution must not be very
	// efficient either. Simply perform DFS from the root until finding the
	// end-node, and push the path to a list of results.
	n := len(graph)
	f := pathFinder{
		results: make([][]int, 0),
	}
	f.path[0] = 0 // Unnecessary, but for clarity
	f.findPaths(graph, 0, 1, n-1)
	return f.results
}

type pathFinder struct {
	path    []int
	results [][]int
}

func (f *pathFinder) findPaths(graph [][]int, curr, pathLen, target int) {
	if curr == target {
		cpy := make([]int, pathLen)
		copy(cpy, f.path[:pathLen])
		f.results = append(f.results, cpy)
		return
	}
	// Try alternatives
	for _, nei := range graph[curr] {
		f.path[pathLen] = nei
		f.findPaths(graph, nei, pathLen+1, target)
	}
}
