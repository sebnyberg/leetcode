package p1066campusbikes2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_assignBikes(t *testing.T) {
	for _, tc := range []struct {
		workers [][]int
		bikes   [][]int
		want    int
	}{
		{[][]int{{0, 0}, {1, 1}, {2, 0}}, [][]int{{1, 0}, {2, 2}, {2, 1}}, 4},
		{[][]int{{0, 0}, {2, 1}}, [][]int{{1, 2}, {3, 3}}, 6},
		{[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, [][]int{{0, 999}, {1, 999}, {2, 999}, {3, 999}, {4, 999}}, 4995},
	} {
		t.Run(fmt.Sprintf("%+v", tc.workers), func(t *testing.T) {
			require.Equal(t, tc.want, assignBikes(tc.workers, tc.bikes))
		})
	}
}

func assignBikes(workers [][]int, bikes [][]int) int {
	n := len(workers)
	m := len(bikes)
	costs := make([]int, 1<<(m+1)-1)
	for i := range costs {
		costs[i] = -1
	}
	// Pre-calculate the distance between workers and bikes
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, m)
		for j := range dist[i] {
			dist[i][j] = abs(workers[i][0]-bikes[j][0]) +
				abs(workers[i][1]-bikes[j][1])
		}
	}
	return int(visit(dist, costs, 0, 0, int(n), int(m)))
}

func visit(dist [][]int, costs []int, bm, idx, n, m int) int {
	if idx == n {
		return 0
	}
	if costs[bm] != -1 {
		return costs[bm]
	}
	minCost := math.MaxInt16
	for j := 0; j < m; j++ {
		var b int = 1 << j
		if bm&b == 0 {
			minCost = min(minCost, dist[idx][j]+visit(dist, costs, bm|b, idx+1, n, m))
		}
	}
	costs[bm] = minCost
	return costs[bm]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
