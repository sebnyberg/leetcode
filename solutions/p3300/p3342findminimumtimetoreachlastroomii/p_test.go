package p3342findminimumtimetoreachlastroomii

import (
	"container/heap"
	"math"
)

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func minTimeToReach(moveTime [][]int) int {
	h := minHeap{&item{0, 0, 0}}
	m := len(moveTime)
	n := len(moveTime[0])
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = 0
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	for len(h) > 0 {
		x := heap.Pop(&h).(*item)
		if x.i == m-1 && x.j == n-1 {
			return x.val
		}
		for _, d := range dirs {
			ii := x.i + d[0]
			jj := x.j + d[1]
			if !ok(ii, jj) {
				continue
			}
			dd := max(dist[x.i][x.j]+1, moveTime[ii][jj]+1)
			dd += (x.i + x.j) & 1
			if dist[ii][jj] <= dd {
				continue
			}
			dist[ii][jj] = dd
			heap.Push(&h, &item{dd, ii, jj})
		}
	}
	return -1
}

type item struct {
	val  int
	i, j int
}

type minHeap []*item

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	it := x.(*item)
	*h = append(*h, it)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
