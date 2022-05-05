package p0882reachablenodesinsubdividedgraph

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reachableNodes(t *testing.T) {
	for _, tc := range []struct {
		edges       [][]int
		maxMoves, n int
		want        int
	}{
		{[][]int{{2, 4, 2}, {3, 4, 5}, {2, 3, 1}, {0, 2, 1}, {0, 3, 5}}, 14, 5, 18},
		{[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}, 6, 3, 13},
		{[][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}}, 10, 4, 23},
		{[][]int{{1, 2, 4}, {1, 4, 6}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}}, 17, 5, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, reachableNodes(tc.edges, tc.maxMoves, tc.n))
		})
	}
}

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	// Approach: use Dijkstra's + count visited nodes along paths.

	// Add edges to an adjacency list
	adj := make([][]path, n)
	for _, edge := range edges {
		a, b, w := edge[0], edge[1], edge[2]
		adj[a] = append(adj[a], path{b, w})
		adj[b] = append(adj[b], path{a, w})
	}

	// Create min-heap
	h := pathHeap{path{0, 0}}

	// Keep track of visited nodes
	visited := make([]bool, n)

	// Coverage of nodes between nodes ("edge nodes") is counted in a map
	// Since there are two edges, these counts need de-duplication at the end
	type edge struct{ u, v int }
	edgeNodes := make(map[edge]int, len(edges)*2)
	dist := make([]int, n)
	for i := 1; i < n; i++ {
		dist[i] = math.MaxInt64
	}

	// Pop paths from min heap
	var res int
	for len(h) > 0 {
		// When visiting a node, it is guaranteed that all edges leaving the node
		// will have the maximum reach possible from that node.
		x := heap.Pop(&h).(path)
		if visited[x.to] {
			continue
		}
		visited[x.to] = true
		res++

		// For each neighbouring node
		for _, a := range adj[x.to] {
			u, v := x.to, a.to
			// Add the maximum number of covered edge nodes to edges map
			edgeNodes[edge{u, v}] = min(maxMoves, x.dist+a.dist) - x.dist

			// Add target node and total distance if smaller than the best distance
			// to the target node so far.
			if d := x.dist + a.dist + 1; d <= maxMoves && d < dist[v] {
				heap.Push(&h, path{v, d})
			}
		}
	}

	// At this point, all visited nodes have been counted, but not edge nodes.
	// Iterate over edges and add covered nodes.
	// Since an edge can be visited twice, the total count may exceed the maximum,
	// in which case the value is truncated.
	for _, e := range edges {
		u, v, numNodes := e[0], e[1], e[2]
		res += min(numNodes, edgeNodes[edge{u, v}]+edgeNodes[edge{v, u}])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type path struct{ to, dist int }

type pathHeap []path

func (h pathHeap) Len() int { return len(h) }
func (h pathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h pathHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}
func (h *pathHeap) Push(x interface{}) {
	*h = append(*h, x.(path))
}
func (h *pathHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
