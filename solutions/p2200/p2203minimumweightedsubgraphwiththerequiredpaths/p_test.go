package p2203minimumweightedsubgraphwiththerequiredpaths

import (
	"container/heap"
	"fmt"
	"leetcode"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumWeight(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		src1  int
		src2  int
		dest  int
		want  int64
	}{
		{
			6,
			leetcode.ParseMatrix("[[0,2,2],[0,5,6],[1,0,3],[1,4,5],[2,1,1],[2,3,3],[2,3,4],[3,4,2],[4,5,1]]"),
			0, 1, 5, 9,
		},
		{
			3,
			leetcode.ParseMatrix("[[0,1,1],[2,1,1]]"),
			0, 1, 2, -1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minimumWeight(tc.n, tc.edges, tc.src1, tc.src2, tc.dest))
		})
	}
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	adjFwd := make([][][2]int, n)
	adjBack := make([][][2]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		adjFwd[u] = append(adjFwd[u], [2]int{v, w})
		adjBack[v] = append(adjBack[v], [2]int{u, w})
	}
	dikstra := func(adj [][][2]int, startIdx int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt64 / 4
		}
		dist[startIdx] = 0
		seen := make([]bool, n)
		h := MinHeap{{startIdx, 0}}
		for len(h) > 0 {
			x := heap.Pop(&h).(item)
			if seen[x.u] {
				continue
			}
			seen[x.u] = true
			for _, near := range adj[x.u] {
				v, w := near[0], near[1]
				if dist[v] <= x.tot+w {
					continue
				}
				dist[v] = x.tot + w
				heap.Push(&h, item{v, x.tot + w})
			}
		}
		return dist
	}

	d1 := dikstra(adjFwd, src1)
	d2 := dikstra(adjFwd, src2)
	d3 := dikstra(adjBack, dest)
	minDist := math.MaxInt64 / 4
	for i := range d1 {
		minDist = min(minDist, d1[i]+d2[i]+d3[i])
	}
	if minDist >= math.MaxInt64/4 {
		return -1
	}
	return int64(minDist)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type item struct {
	u, tot int
}

type MinHeap []item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].tot < h[j].tot
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
