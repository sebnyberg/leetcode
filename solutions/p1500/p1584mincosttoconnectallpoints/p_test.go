package p1584mincosttoconnectallpoints

import (
	"fmt"
	"sort"
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

	// Calculate and sort distances between all pairs of points
	pairDistances := make([]pointDistance, 0)
	for i := range points {
		for j := i + 1; j < n; j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			pairDistances = append(pairDistances, pointDistance{
				i, j, abs(x2-x1) + abs(y2-y1),
			})
		}
	}
	sort.Slice(pairDistances, func(i, j int) bool {
		return pairDistances[i].dist < pairDistances[j].dist
	})

	// Add distances between pairs of points until all
	// points are in the same set
	dsu := NewDSU(n)
	dist := 0
	for _, pair := range pairDistances {
		if dsu.find(pair.i) == dsu.find(pair.j) {
			continue
		}
		dsu.union(pair.i, pair.j)
		dist += pair.dist
	}

	return dist
}

type pointDistance struct {
	i, j int
	dist int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b int) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		if d.size[a] < d.size[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
	}
}
