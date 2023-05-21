package p1334findthecitywiththesmallestnumberofneigborsatathresholddistance

import (
	"container/heap"
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// We can use a form of modified dijkstra's where we aim to visit as many
	// nodes as we can below the threshold
	dist := make([]int, n)
	type edge struct {
		to int
		w  int
	}
	adj := make([][]edge, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		w := e[2]
		adj[a] = append(adj[a], edge{b, w})
		adj[b] = append(adj[b], edge{a, w})
	}
	h := make(visitHeap, 0, n)
	countReachable := func(start int) int {
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		h = append(h, visit{start, 0})
		dist[start] = 0
		var count int
		for len(h) > 0 {
			x := heap.Pop(&h).(visit)
			if dist[x.j] < x.dist {
				continue
			}
			for _, e := range adj[x.j] {
				d := x.dist + e.w
				if d > distanceThreshold || dist[e.to] <= d {
					continue
				}
				if dist[e.to] == math.MaxInt32 {
					count++
				}
				dist[e.to] = d
				heap.Push(&h, visit{e.to, d})
			}
		}
		return count
	}
	minCount := math.MaxInt32
	var maxIdx int
	for i := 0; i < n; i++ {
		count := countReachable(i)
		if count <= minCount {
			minCount = count
			maxIdx = i
		}
	}
	return maxIdx
}

type visit struct {
	j    int
	dist int
}

type visitHeap []visit

func (h visitHeap) Len() int { return len(h) }
func (h visitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h visitHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}
func (h *visitHeap) Push(x interface{}) {
	el := x.(visit)
	*h = append(*h, el)
}
func (h *visitHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
