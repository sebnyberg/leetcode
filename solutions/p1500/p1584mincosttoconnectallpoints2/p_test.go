package p1584mincosttoconnectallpoints

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCostConnectPoints(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int
	}{
		{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
		{[][]int{{3, 12}, {-2, 5}, {-4, 1}}, 18},
		{[][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}}, 4},
		{[][]int{{-1000000, -1000000}, {1000000, 1000000}}, 4000000},
		{[][]int{{0, 0}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, minCostConnectPoints(tc.points))
		})
	}
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dist1d := func(p1, p2 []int) int {
		return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
	}
	var visited [16]uint64
	var dist [1001]int
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	curr := 0
	dist[0] = 0
	var res int
	// For each point
	for nvisited := 1; nvisited <= n; nvisited++ {
		visited[curr/64] |= (1 << (curr % 64))
		res += dist[curr]

		// Visit all neighbouring points
		minDist, nextIdx := math.MaxInt32, -1
		for nei := 0; nei < n; nei++ {
			// If both points are visited, do nothing
			if visited[nei/64]&(1<<(nei%64)) > 0 {
				continue
			}
			// If shortest distance to the target point, update
			d := dist1d(points[curr], points[nei])
			if d < dist[nei] {
				dist[nei] = d
			}

			// If shortest distance out of all points,
			// mark as next place to visit
			if dist[nei] < minDist {
				minDist = dist[nei]
				nextIdx = nei
			}
		}
		curr = nextIdx
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
