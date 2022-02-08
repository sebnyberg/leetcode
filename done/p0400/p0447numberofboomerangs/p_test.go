package p0447numberofboomerangs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfBoomerangs(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int
	}{
		{[][]int{{0, 0}, {1, 0}, {2, 0}}, 2},
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 2},
		{[][]int{{1, 1}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfBoomerangs(tc.points))
		})
	}
}

func numberOfBoomerangs(points [][]int) int {
	// Since all points are unique, we can select arbitrary points, check distance
	// to each other point, and check if there is a point on the opposite side.
	// pts := make(map[[2]int]struct{}, len(points))
	// for _, p := range points {
	// 	pts[[2]int{p[0], p[1]}] = struct{}{}
	// }
	dist := func(i, j int) int {
		p1 := points[i]
		p2 := points[j]
		dx := p2[0] - p1[0]
		dy := p2[1] - p1[1]
		return dx*dx + dy*dy
	}

	var res int
	dists := make(map[int]int, len(points))
	for i := 0; i < len(points); i++ {
		for k := range dists {
			dists[k] = 0
		}
		for j := 0; j < len(points); j++ {
			dists[dist(i, j)]++
		}
		for _, val := range dists {
			res += val * (val - 1)
		}
	}
	return res
}
