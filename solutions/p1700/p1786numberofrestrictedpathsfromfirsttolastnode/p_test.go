package p1786numberofrestrictedpathsfromfirsttolastnode

import (
	"container/heap"
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
		{7, [][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}}, 1},
		{5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}, 3},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.n, tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, countRestrictedPaths(tc.n, tc.edges))
		})
	}
}

type weightedEdge struct {
	weight int
	node   int
}

func countRestrictedPaths(n int, edges [][]int) int {
	adj := make([][]weightedEdge, n+1)
	for _, edge := range edges {
		a, b, weight := edge[0], edge[1], edge[2]
		adj[a] = append(adj[a], weightedEdge{weight, b})
		adj[b] = append(adj[b], weightedEdge{weight, a})
	}

	// Dijkstra (distance from n to a node)
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[n] = 0
	h := WeightedEdgeHeap{weightedEdge{0, n}} // distance to n is zero
	heap.Init(&h)
	counts := make([]int, n+1)
	counts[n] = 1
	visited := make([]bool, n+1)
	for h.Len() > 0 {
		cur := heap.Pop(&h).(weightedEdge)
		if visited[cur.node] {
			continue
		}
		visited[cur.node] = true
		for _, near := range adj[cur.node] {
			if !visited[near.node] {
				if dist[near.node] > dist[cur.node] {
					counts[near.node] += counts[cur.node]
					counts[near.node] %= mod
				}
				if d := dist[cur.node] + near.weight; d < dist[near.node] {
					dist[near.node] = d
					heap.Push(&h, weightedEdge{d, near.node})
				}
			}
		}
	}

	return counts[1] % mod
}

const mod = 1000000007

type WeightedEdgeHeap []weightedEdge

func (h WeightedEdgeHeap) Len() int { return len(h) }
func (h WeightedEdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h WeightedEdgeHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}
func (h *WeightedEdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(weightedEdge))
}
func (h *WeightedEdgeHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
