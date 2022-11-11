package p0963minimumarearectangleii

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minAreaFreeRect(t *testing.T) {
	for i, tc := range []struct {
		points [][]int
		want   float64
	}{
		{
			leetcode.ParseMatrix("[[3,1],[1,1],[0,1],[2,1],[3,3],[3,2],[0,2],[2,3]]"),
			2,
		},
		{
			leetcode.ParseMatrix("[[1,2],[2,1],[1,0],[0,1]]"),
			2,
		},
		{
			leetcode.ParseMatrix("[[0,1],[2,1],[1,1],[1,0],[2,0]]"),
			1,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.InEpsilon(t, tc.want, minAreaFreeRect(tc.points), 0.0001)
		})
	}
}

func minAreaFreeRect(points [][]int) float64 {
	m := make(map[[2]int]struct{})
	for _, p := range points {
		m[[2]int{p[0], p[1]}] = struct{}{}
	}
	exists := func(i, j int) bool {
		_, ok := m[[2]int{i, j}]
		return ok
	}

	// For each point, pick two other points. Draw a vector to those points.
	// If the angle is 90 degrees, calculate where the fourth point should be.
	res := math.MaxFloat64
	for i := range points {
		for j := range points {
			for k := range points {
				if i == j || j == k || i == k {
					continue
				}
				mid := points[i]
				p := points[j]
				q := points[k]
				dxp := p[0] - mid[0]
				dyp := p[1] - mid[1]
				dxq := q[0] - mid[0]
				dyq := q[1] - mid[1]
				dot := dxp*dxq + dyp*dyq
				if dot != 0 {
					continue
				}
				xx := mid[0] + dxp + dxq
				yy := mid[1] + dyp + dyq
				if exists(xx, yy) {
					area := math.Sqrt(float64(dxp*dxp+dyp*dyp)) *
						math.Sqrt(float64(dxq*dxq+dyq*dyq))
					if area < res {
						res = area
					}
				}
			}
		}
	}

	if res < math.MaxFloat64 {
		return res
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
