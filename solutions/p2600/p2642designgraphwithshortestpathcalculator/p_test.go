package p2642designgraphwithshortestpathcalculator

import (
	"container/heap"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	edges := leetcode.ParseMatrix("[[0,2,5],[0,1,2],[1,2,1],[3,0,3]]")
	c := Constructor(4, edges)
	res := c.ShortestPath(3, 2)
	require.Equal(t, res, 6)
	res = c.ShortestPath(0, 3)
	require.Equal(t, res, -1)
	c.AddEdge([]int{1, 3, 4})
	res = c.ShortestPath(0, 3)
	require.Equal(t, res, 6)
}

type edge struct {
	to int
	w  int
}

type visit struct {
	u    int
	dist int
}

type Graph struct {
	adj  [][]edge
	dist []int
	h    minHeap
}

func Constructor(n int, edges [][]int) Graph {
	var g Graph
	g.adj = make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.adj[u] = append(g.adj[u], edge{v, w})
	}
	g.dist = make([]int, n)
	g.h = make(minHeap, 0, n)
	return g
}

func (this *Graph) AddEdge(e []int) {
	u, v, w := e[0], e[1], e[2]
	this.adj[u] = append(this.adj[u], edge{v, w})
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	for i := range this.dist {
		this.dist[i] = math.MaxInt64
	}
	this.h = this.h[:0]
	heap.Push(&this.h, visit{node1, 0})
	for len(this.h) > 0 {
		x := heap.Pop(&this.h).(visit)
		if x.u == node2 {
			return x.dist
		}
		if this.dist[x.u] < x.dist {
			continue
		}
		for _, y := range this.adj[x.u] {
			dd := x.dist + y.w
			if this.dist[y.to] <= dd {
				continue
			}
			this.dist[y.to] = dd
			heap.Push(&this.h, visit{y.to, dd})
		}
	}
	return -1
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
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
