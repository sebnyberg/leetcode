package p1786numberofrestrictedpathsfromfirsttolastnode

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countRestrictedPaths(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		want  int
	}{
		// {5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}, 3},
		// {7, [][]int{{1, 5, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}}, 3},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.n, tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, countRestrictedPaths(tc.n, tc.edges))
		})
	}
}

// TODO
type weightedEdge struct {
	weight int
	to     int
}

func countRestrictedPaths(n int, edges [][]int) int {
	adj := make([][]*weightedEdge, n+1)
	for _, edge := range edges {
		a, b, weight := edge[0], edge[1], edge[2]
		adj[a] = append(adj[a], &weightedEdge{weight, b})
		adj[b] = append(adj[b], &weightedEdge{weight, a})
	}

	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -1
	}
	for i := 1; i < n; i++ {
		findShortestPath(mem, i, n, adj)
	}

	nrestricted := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if mem[i] > mem[j] {
				nrestricted++
			}
		}
	}

	return nrestricted
}

func findShortestPath(mem []int, i, n int, adjWeights [][]*weightedEdge) int {
	if mem[i] != -1 {
		return mem[i]
	}
	if i == n {
		return 0
	}
	minSum := math.MaxInt32
	for _, edge := range adjWeights[i] {
		if edge.to <= i {
			continue
		}
		minSum = min(minSum, edge.weight+findShortestPath(mem, edge.to, n, adjWeights))
	}
	mem[i] = minSum
	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
