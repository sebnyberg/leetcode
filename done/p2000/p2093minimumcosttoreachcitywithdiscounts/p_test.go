package p2093minimumcosttoreachcitywithdiscounts

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCost(t *testing.T) {
	for _, tc := range []struct {
		n         int
		highways  [][]int
		discounts int
		want      int
	}{
		{5, [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 3}, {3, 4, 2}}, 1, 9},
		{4, [][]int{{1, 3, 17}, {1, 2, 7}, {3, 2, 5}, {0, 1, 6}, {3, 0, 20}}, 20, 8},
		{4, [][]int{{0, 1, 3}}, 0, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minimumCost(tc.n, tc.highways, tc.discounts))
		})
	}
}

func minimumCost(n int, highways [][]int, discounts int) int {
	// Use modified Dijkstra's to visit all nodes, lowest distance first.
	// The modification is that each possible discount has its own distance
	// vector.
	adj := make([][]int, n)
	tolls := make([][]int, n)
	for _, h := range highways {
		u, v, toll := h[0], h[1], h[2]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		tolls[u] = append(tolls[u], toll)
		tolls[v] = append(tolls[v], toll)
	}

	dists := make([][]int, discounts+1)
	for i := range dists {
		dists[i] = make([]int, n)
		for j := range dists[i] {
			dists[i][j] = math.MaxInt32
		}
	}
	dists[discounts][0] = 0

	h := VisitHeap{visit{0, 0, discounts}}
	for len(h) > 0 {
		x := heap.Pop(&h).(visit)
		if x.at == n-1 { // Reached the end
			return x.dist
		}
		for i, near := range adj[x.at] {
			// Try with and without discount
			discountDist := x.dist + tolls[x.at][i]/2
			if x.discounts > 0 && discountDist < dists[x.discounts-1][near] {
				// Remove sub-optimal paths due to initial math.MaxInt32 distance
				for d := x.discounts - 1; d >= 0 && dists[d][near] > discountDist; d-- {
					dists[d][near] = discountDist
				}
				heap.Push(&h, visit{near, discountDist, x.discounts - 1})
			}
			noDiscountDist := x.dist + tolls[x.at][i]
			if noDiscountDist < dists[x.discounts][near] {
				// Remove sub-optimal paths due to initial math.MaxInt32 distance
				for d := x.discounts; d >= 0 && dists[d][near] > noDiscountDist; d-- {
					dists[d][near] = noDiscountDist
				}
				heap.Push(&h, visit{near, noDiscountDist, x.discounts})
			}
		}
	}
	return -1
}

type visit struct {
	at        int
	dist      int
	discounts int
}

type VisitHeap []visit

func (h VisitHeap) Len() int { return len(h) }
func (h VisitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h VisitHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
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
