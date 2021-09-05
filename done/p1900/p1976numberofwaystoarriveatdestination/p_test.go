package p1976numberofwaystoarriveatdestination

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPaths(t *testing.T) {
	for _, tc := range []struct {
		n     int
		roads [][]int
		want  int
	}{
		{
			7,
			[][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}},
			4,
		},
		{2, [][]int{{1, 0, 10}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countPaths(tc.n, tc.roads))
		})
	}
}

const mod = 1_000_000_007

func countPaths(n int, roads [][]int) int {
	adj := make([][]path, n)
	for _, road := range roads {
		a, b, t := road[0], road[1], road[2]
		adj[a] = append(adj[a], path{a, b, t})
		adj[b] = append(adj[b], path{b, a, t})
	}

	// Since travel time is non-negative, we can use Dijkstra's algorithm
	// However, the stopping condition is not to reach the end node, but to
	// reach the end node with a time greater than the minimum time.
	time := make([]int, n+1)
	nVisits := make([]int, n+1)
	for i := range time {
		time[i] = math.MaxInt64
	}
	time[n] = 0
	nVisits[n] = 1
	// visited := make([]bool, n)
	h := pathHeap{path{n, 0, 0}}
	for len(h) > 0 {
		// Pop all paths which share the same distance
		x := heap.Pop(&h).(path)
		if nVisits[x.to] > 0 {
			continue
		}
		curTime := x.time
		// Add all paths which do not already have a recorded path
		tovisit := make(map[int][]path)
		tovisit[x.to] = append(tovisit[x.to], x)
		for len(h) > 0 && h[0].time == curTime {
			a := heap.Pop(&h).(path)
			if nVisits[a.to] > 0 {
				continue
			}
			tovisit[a.to] = append(tovisit[a.to], a)
		}
		// Increment number of ways to reach a node in min time by adding
		// the number of ways to reach the previous node.
		for u, paths := range tovisit {
			time[u] = curTime
			for _, path := range paths {
				nVisits[u] += nVisits[path.from]
				nVisits[u] %= mod
			}
			for _, nei := range adj[u] {
				if curTime+nei.time <= time[nei.to] {
					heap.Push(&h, path{u, nei.to, curTime + nei.time})
				}
			}
		}
		if nVisits[n-1] != 0 {
			return nVisits[n-1]
		}
	}
	return -1
}

type path struct {
	from int
	to   int
	time int
}

type pathHeap []path

func (h pathHeap) Len() int { return len(h) }
func (h pathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h pathHeap) Less(i, j int) bool {
	return h[i].time < h[j].time
}
func (h *pathHeap) Push(x interface{}) {
	*h = append(*h, x.(path))
}
func (h *pathHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
