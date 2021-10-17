package p2045secondminimumtoreachdestination

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_secondMinimum(t *testing.T) {
	for _, tc := range []struct {
		n            int
		edges        [][]int
		time, change int
		want         int
	}{
		{6, [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {5, 4}, {4, 6}}, 3, 100, 12},
		{5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5, 13},
		{2, [][]int{{1, 2}}, 3, 2, 11},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, secondMinimum(tc.n, tc.edges, tc.time, tc.change))
		})
	}
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// Perform Dijkstra's with the modification that each node can be visited
	// twice, and there are two (unique) min times.
	adj := make([][]uint16, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adj[a] = append(adj[a], uint16(b))
		adj[b] = append(adj[b], uint16(a))
	}

	minTimes := make([][2]uint32, n+1)
	for i := range minTimes {
		minTimes[i] = [2]uint32{math.MaxInt32, math.MaxInt32}
	}

	h := DistanceHeap{weightedEdge{0, 1}}
	for len(h) > 0 {
		e := heap.Pop(&h).(weightedEdge)
		u, t := e.node, e.arrivalTime
		if t == minTimes[u][0] || t >= minTimes[u][1] {
			continue
		}

		if t < minTimes[u][0] {
			minTimes[u][0] = t
		} else {
			minTimes[u][1] = t
		}

		// If light is currently red, wait for a green signal.
		d := t / uint32(change)
		if isGreen := d%2 == 0; !isGreen {
			nextGreen := (d + 1) * uint32(change)
			t = max(t, nextGreen)
		}
		for _, nei := range adj[u] {
			if t+uint32(time) >= minTimes[nei][1] {
				continue
			}
			heap.Push(&h, weightedEdge{
				arrivalTime: t + uint32(time),
				node:        nei,
			})
		}
	}
	return int(minTimes[n][1])
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

type weightedEdge struct {
	arrivalTime uint32
	node        uint16
}

type DistanceHeap []weightedEdge

func (h DistanceHeap) Len() int { return len(h) }
func (h DistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h DistanceHeap) Less(i, j int) bool {
	return h[i].arrivalTime < h[j].arrivalTime
}
func (h *DistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(weightedEdge))
}
func (h *DistanceHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
