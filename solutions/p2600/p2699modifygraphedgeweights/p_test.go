package p2699modifygraphedgeweights

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_modifiedGraphEdges(t *testing.T) {
	for i, tc := range []struct {
		n           int
		edges       [][]int
		source      int
		destination int
		target      int
		want        [][]int
	}{
		{
			100,
			leetcode.ParseMatrix("[[0,1,10000000],[1,2,10000000],[2,3,10000000],[3,4,10000000],[4,5,10000000],[5,6,10000000],[6,7,10000000],[7,8,10000000],[8,9,10000000],[9,10,10000000],[10,11,10000000],[11,12,10000000],[12,13,10000000],[13,14,10000000],[14,15,10000000],[15,16,10000000],[16,17,10000000],[17,18,10000000],[18,19,10000000],[19,20,10000000],[20,21,10000000],[21,22,10000000],[22,23,10000000],[23,24,10000000],[24,25,10000000],[25,26,10000000],[26,27,10000000],[27,28,10000000],[28,29,10000000],[29,30,10000000],[30,31,10000000],[31,32,10000000],[32,33,10000000],[33,34,10000000],[34,35,10000000],[35,36,10000000],[36,37,10000000],[37,38,10000000],[38,39,10000000],[39,40,10000000],[40,41,10000000],[41,42,10000000],[42,43,10000000],[43,44,10000000],[44,45,10000000],[45,46,10000000],[46,47,10000000],[47,48,10000000],[48,49,10000000],[49,50,10000000],[50,51,10000000],[51,52,10000000],[52,53,10000000],[53,54,10000000],[54,55,10000000],[55,56,10000000],[56,57,10000000],[57,58,10000000],[58,59,10000000],[59,60,10000000],[60,61,10000000],[61,62,10000000],[62,63,10000000],[63,64,10000000],[64,65,10000000],[65,66,10000000],[66,67,10000000],[67,68,10000000],[68,69,10000000],[69,70,10000000],[70,71,10000000],[71,72,10000000],[72,73,10000000],[73,74,10000000],[74,75,10000000],[75,76,10000000],[76,77,10000000],[77,78,10000000],[78,79,10000000],[79,80,10000000],[80,81,10000000],[81,82,10000000],[82,83,10000000],[83,84,10000000],[84,85,10000000],[85,86,10000000],[86,87,10000000],[87,88,10000000],[88,89,10000000],[89,90,10000000],[90,91,10000000],[91,92,10000000],[92,93,10000000],[93,94,10000000],[94,95,10000000],[95,96,10000000],[96,97,10000000],[97,98,10000000],[98,99,10000000],[0,99,-1]]"),
			0, 99, 990000000,
			leetcode.ParseMatrix("[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]"),
		},
		{
			4,
			leetcode.ParseMatrix("[[0,1,5],[1,2,7],[2,3,7],[3,1,9],[3,0,-1],[0,2,-1]]"),
			2, 3, 7,
			leetcode.ParseMatrix("[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]"),
		},
		{
			4,
			leetcode.ParseMatrix("[[1,0,4],[1,2,3],[2,3,5],[0,3,-1]]"),
			0, 2, 6,
			leetcode.ParseMatrix("[[1,0,4],[1,2,3],[2,3,5],[0,3,1]]"),
		},
		{
			5,
			leetcode.ParseMatrix("[[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]]"),
			0, 1, 5,
			leetcode.ParseMatrix("[[4,1,1],[2,0,1],[0,3,3],[4,3,1]]"),
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, modifiedGraphEdges(tc.n, tc.edges, tc.source, tc.destination, tc.target))
		})
	}
}

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	// In my head this makes sense but we will see if it works.
	//
	// For a solution to exist, the shortest path from source to destination
	// must include at least one flexible edge. If we consider all unweighted
	// edges to be math.MaxInt32 and the total distance from source to dest is
	// <= target, then there is no solution.
	//
	// There is possibly another case where, even if the weighted edges were all
	// 1, the shortest distance to the destination would still be > target. Then
	// it does not matter what we do with the flexible edges.
	//
	// There is an iterative method where we simply run the shortest path
	// algorithm over and over, adding weights to flexible edges until the total
	// sum equals the target. Maybe start by having all edges be 1?
	//
	type weightedEdge struct {
		a        int
		b        int
		w        int
		flexible bool
	}
	flexible := make([]*weightedEdge, 0, n)
	adj := make([][]*weightedEdge, n)
	adjMatrix := make([][]*weightedEdge, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]*weightedEdge, n)
	}
	for _, e := range edges {
		a := e[0]
		b := e[1]
		w := e[2]
		if a > b {
			a, b = b, a
		}
		edge := &weightedEdge{a, b, w, w < 0}
		if w < 0 {
			flexible = append(flexible, edge)
		}
		adj[a] = append(adj[a], edge)
		adj[b] = append(adj[b], edge)
		adjMatrix[a][b] = edge
		adjMatrix[b][a] = edge
	}

	prev := make([]int, n)
	dist := make([]int, n)
	h := make(minHeap, 0, n)
	path := make([]int, 0, n)

	// findShortest returns the shortest path from the source to the
	// destination and its total distance
	findShortest := func() ([]int, int) {
		for i := range prev {
			prev[i] = i
		}
		for i := range dist {
			dist[i] = math.MaxInt64 >> 2
		}
		dist[source] = 0
		h = append(h[:0], visit{source, source, 0})
		for len(h) > 0 {
			x := heap.Pop(&h).(visit)
			if x.dist > dist[x.to] {
				if prev[x.to] == x.from {
					panic("wut!!")
				}
				continue
			}
			for _, nei := range adj[x.to] {
				d := dist[x.to] + nei.w
				next := nei.b
				if nei.b == x.to {
					next = nei.a
				}
				if d >= dist[next] {
					continue
				}
				dist[next] = d
				heap.Push(&h, visit{x.to, next, d})
				prev[next] = x.to
			}
		}
		curr := destination
		path = path[:0]
		path = append(path, curr)
		for curr != source {
			curr = prev[curr]
			path = append(path, curr)
		}
		return path, dist[destination]
	}
	_ = findShortest

	// Set flexible edges to math.MaxInt32. If there's a path < parget then
	// there is no solution
	var shortestPath []int
	var shortestPathDist int
	for _, e := range flexible {
		e.w = 2e9
	}
	shortestPath, shortestPathDist = findShortest()
	if shortestPathDist < target {
		return [][]int{}
	}
	if shortestPathDist == target {
		for i, e := range edges {
			edges[i][2] = adjMatrix[e[0]][e[1]].w
		}
		return edges
	}
	if len(flexible) == 0 {
		return [][]int{}
	}

	for _, e := range flexible {
		e.w = 1
	}
	for {
		shortestPath, shortestPathDist = findShortest()
		if shortestPathDist == target {
			break
		}
		if shortestPathDist > target {
			return [][]int{}
		}
		// Modify first flexible edge in the path to add the missing value
		// toward target.
		missing := target - shortestPathDist
		for i := len(shortestPath) - 1; i >= 1; i-- {
			a := shortestPath[i]
			b := shortestPath[i-1]
			e := adjMatrix[a][b]
			if e.flexible {
				e.w = min(2*1e9, 1+missing)
				break
			}
		}
	}
	for i, e := range edges {
		edges[i][2] = adjMatrix[e[0]][e[1]].w
	}
	return edges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type visit struct {
	from int
	to   int
	dist int
}

type minHeap []visit

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}
func (h *minHeap) Push(x interface{}) {
	el := x.(visit)
	// el.heapIdx = len(*h)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	// el = nil
	*h = (*h)[:n-1]
	return el
}
