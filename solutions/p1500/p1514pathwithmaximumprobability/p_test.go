package p1514pathwithmaximumprobability

import (
	"container/heap"
)

type weightedEdge struct {
	prob float64
	node int
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	adj := make([][]weightedEdge, n+1)
	for i, edge := range edges {
		a, b := edge[0], edge[1]
		weight := succProb[i]
		adj[a] = append(adj[a], weightedEdge{weight, b})
		adj[b] = append(adj[b], weightedEdge{weight, a})
	}

	// Dijkstra (distance from n to a node)
	probabilities := make([]float64, n+1)
	for i := range probabilities {
		probabilities[i] = 0
	}
	probabilities[start] = 1
	h := WeightedEdgeHeap{weightedEdge{1, start}} // distance to start is 1
	heap.Init(&h)
	for h.Len() > 0 {
		cur := heap.Pop(&h).(weightedEdge)
		if probabilities[cur.node] > cur.prob {
			continue
		}
		for _, near := range adj[cur.node] {
			d := cur.prob * near.prob
			if probabilities[near.node] >= d {
				continue
			}
			probabilities[near.node] = d
			heap.Push(&h, weightedEdge{d, near.node})
		}
	}
	return probabilities[end]
}

type WeightedEdgeHeap []weightedEdge

func (h WeightedEdgeHeap) Len() int { return len(h) }
func (h WeightedEdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h WeightedEdgeHeap) Less(i, j int) bool {
	return h[i].prob > h[j].prob
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
