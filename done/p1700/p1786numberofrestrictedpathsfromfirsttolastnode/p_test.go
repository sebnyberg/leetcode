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
		{5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}, 3},
		// {7, [][]int{{1, 5, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}}, 3},
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
	adj := make([][]*weightedEdge, n+1)
	for _, edge := range edges {
		a, b, weight := edge[0], edge[1], edge[2]
		adj[a] = append(adj[a], &weightedEdge{weight, b})
		adj[b] = append(adj[b], &weightedEdge{weight, a})
	}

	// Dijkstra (distance from n to a node)
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[n] = 0
	h := WeightedEdgeHeap{&weightedEdge{0, n}}
	heap.Init(&h)
	for h.Len() > 0 {
		edge := heap.Pop(&h).(*weightedEdge)
		if edge.weight != dist[edge.node] { // not shortest path
			continue
		}
		for _, near := range adj[edge.node] {
			if d := dist[edge.node] + near.weight; d < dist[near.node] {
				dist[near.node] = d
				heap.Push(&h, &weightedEdge{dist[near.node], near.node})
			}
		}
	}

	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -1
	}

	return findRestricted(mem, dist, adj, n, 1)
}

const mod = 1000000007

func findRestricted(mem []int, dist []int, adj [][]*weightedEdge, n, i int) int {
	if mem[i] != -1 {
		return mem[i]
	}
	if i == n {
		return 1
	}
	res := 0
	for _, near := range adj[i] {
		if dist[near.node] > dist[i] {
			res += findRestricted(mem, dist, adj, n, near.node)
			res %= mod
		}
	}
	mem[i] = res
	return res
}

type WeightedEdgeHeap []*weightedEdge

func (h WeightedEdgeHeap) Len() int { return len(h) }
func (h WeightedEdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h WeightedEdgeHeap) Less(i, j int) bool {
	return h[i].weight < h[i].weight
}
func (h *WeightedEdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(*weightedEdge))
}
func (h *WeightedEdgeHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
