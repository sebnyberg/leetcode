package p3607powergridmaintenance

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_processQueries(t *testing.T) {
	for _, tc := range []struct {
		c           int
		connections [][]int
		queries     [][]int
		want        []int
	}{
		{
			3,
			leetcode.ParseMatrix("[[3,2],[1,3],[2,1]]"),
			leetcode.ParseMatrix("[[2,2],[1,2],[1,2],[1,3],[1,1],[1,3],[1,1],[1,1],[2,1],[1,1],[2,3],[2,3],[2,3],[2,1],[2,1],[2,1],[1,1],[1,1],[1,2],[1,2],[2,1],[2,1],[2,2],[1,2],[1,1]]"),
			[]int{1, 1, 3, 1, 3, 1, 1, 3, -1, -1, -1, -1, -1, -1},
		},
		{5, leetcode.ParseMatrix("[[1,2],[2,3],[3,4],[4,5]]"), leetcode.ParseMatrix("[[1,3],[2,1],[1,1],[2,2],[1,2]]"), []int{3, 2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.c), func(t *testing.T) {
			require.Equal(t, tc.want, processQueries(tc.c, tc.connections, tc.queries))
		})
	}
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	// First, process connections and perform union-find to partition the stations
	parent := make([]int, c+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		ra := parent[x]
		if ra != x {
			ra = find(ra)
		}
		parent[x] = ra
		return ra
	}

	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra > rb {
			ra, rb = rb, ra
		}
		parent[rb] = ra
	}

	for _, cc := range connections {
		union(cc[0], cc[1])
	}

	clusters := make(map[int]*minHeap)
	stations := make([]*station, c+1)
	for i := range stations {
		stations[i] = &station{station: i, offline: false}
	}
	for k := 1; k <= c; k++ {
		ra := find(k)
		if clusters[ra] == nil {
			clusters[ra] = &minHeap{}
		}
		heap.Push(clusters[ra], stations[k])
	}

	var res []int
	for i, q := range queries {
		_ = i
		op := q[0]
		stationIdx := q[1]
		clusterIndex := find(stationIdx)
		switch {
		case op == 2:
			stations[stationIdx].offline = true
			heap.Fix(clusters[clusterIndex], stations[stationIdx].idx)
		// health check
		case !stations[stationIdx].offline:
			res = append(res, stationIdx)
		// station is offline
		default:
			x := (*clusters[clusterIndex])[0]
			if x.offline { // "best" station is offline
				res = append(res, -1)
			} else {
				res = append(res, x.station)
			}
		}
	}

	return res
}

type station struct {
	station int
	offline bool
	idx     int
}

type minHeap []*station

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h minHeap) Less(i, j int) bool {
	a := h[i]
	b := h[j]
	if a.offline != b.offline {
		return !a.offline
	}
	return a.station < b.station
}
func (h *minHeap) Push(x interface{}) {
	it := x.(*station)
	it.idx = len(*h)
	*h = append(*h, it)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
