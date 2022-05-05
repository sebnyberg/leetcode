package p1928minimumcosttoreachdestinationintime

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for _, tc := range []struct {
		maxTime     int
		edges       [][]int
		passingFees []int
		want        int
	}{
		{30, [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}, []int{5, 1, 2, 20, 20, 3}, 11},
		{29, [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}, []int{5, 1, 2, 20, 20, 3}, 48},
		{25, [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}}, []int{5, 1, 2, 20, 20, 3}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.maxTime), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.maxTime, tc.edges, tc.passingFees))
		})
	}
}

type edge struct {
	to, time int
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// Dijkstra's:
	// Keep a list of minimum cost to reach a given node.
	// If a smaller cost path, or a path which has smaller distance is found
	// to a node, replace it in the list of "distances".
	n := len(passingFees)
	adj := make([][]edge, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], edge{b, e[2]})
		adj[b] = append(adj[b], edge{a, e[2]})
	}

	time := make([]int, n)
	cost := make([]int, n)
	for i := range time {
		time[i] = math.MaxInt32
		cost[i] = math.MaxInt32
	}
	cost[0] = passingFees[0]
	time[0] = 0

	h := tripHeap{trip{0, passingFees[0], 0}}
	for len(h) > 0 {
		x := heap.Pop(&h).(trip)
		if x.at == n-1 {
			return x.cost
		}
		for _, nei := range adj[x.at] {
			c := x.cost + passingFees[nei.to]
			t := x.time + nei.time
			if t > maxTime {
				continue
			}
			if c < cost[nei.to] {
				cost[nei.to] = c
				time[nei.to] = t
				heap.Push(&h, trip{nei.to, c, t})
			} else if t < time[nei.to] {
				time[nei.to] = t
				heap.Push(&h, trip{nei.to, c, t})
			}
		}
	}

	if cost[n-1] == math.MaxInt32 {
		return -1
	}
	return cost[n-1]
}

type trip struct {
	at   int
	cost int
	time int
}

type tripHeap []trip

func (h tripHeap) Len() int { return len(h) }
func (h tripHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h tripHeap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].time < h[j].time
	}
	return h[i].cost < h[j].cost
}
func (h *tripHeap) Push(x interface{}) {
	*h = append(*h, x.(trip))
}
func (h *tripHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
