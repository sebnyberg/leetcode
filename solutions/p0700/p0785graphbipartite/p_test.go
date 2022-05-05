package p0785graphbipartite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isBipartite(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		want  bool
	}{
		{[][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}, false},
		{[][]int{{4}, {}, {4}, {4}, {0, 2, 3}}, true},
		{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			require.Equal(t, tc.want, isBipartite(tc.graph))
		})
	}
}

func isBipartite(graph [][]int) bool {
	a := make(map[int]bool)
	b := make(map[int]bool)
	visited := make(map[int]bool)
	for i := range graph {
		if !visited[i] {
			if !dfs(i, 0, graph, a, b, visited) {
				return false
			}
		}
	}
	return true
}

func dfs(i int, depth int, graph [][]int, a, b, visited map[int]bool) bool {
	isA := depth%2 == 0
	if isA {
		a[i] = true
	} else {
		b[i] = true
	}
	visited[i] = true

	for _, edge := range graph[i] {
		if visited[edge] {
			if isA && !b[edge] {
				return false
			} else if !isA && !a[edge] {
				return false
			}
		} else {
			if !dfs(edge, depth+1, graph, a, b, visited) {
				return false
			}
		}
	}
	return true
}
