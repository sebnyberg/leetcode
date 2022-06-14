package p0743networkdelaytime

import (
	"container/heap"
	"math"
)

type visit struct {
	u, totalDist int
}

type edge struct {
	v, w int
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][]edge, n+1)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		adj[u] = append(adj[u], edge{v, w})
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	h := VisitHeap{visit{k, 0}}
	seen := make([]bool, n+1)
	var seenCount int
	for len(h) > 0 {
		x := heap.Pop(&h).(visit)
		if dist[x.u] < x.totalDist {
			continue
		}
		seen[x.u] = true
		seenCount++
		for _, nei := range adj[x.u] {
			d := x.totalDist + nei.w
			if d >= dist[nei.v] {
				continue
			}
			dist[nei.v] = d
			heap.Push(&h, visit{nei.v, d})
		}
	}
	if seenCount != n {
		return -1
	}
	var maxDist int
	for _, d := range dist[1:] {
		if d > maxDist {
			maxDist = d
		}
	}
	return maxDist
}

type VisitHeap []visit

func (h VisitHeap) Len() int { return len(h) }
func (h VisitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h VisitHeap) Less(i, j int) bool {
	return h[i].totalDist < h[j].totalDist
}
func (h *VisitHeap) Push(x interface{}) {
	*h = append(*h, x.(visit))
}
func (h *VisitHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
