package p0469convexpolygon

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isConvex(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   bool
	}{
		{leetcode.ParseMatrix("[[0,0],[0,5],[5,5],[5,0]]"), true},
		{leetcode.ParseMatrix("[[0,0],[0,10],[10,10],[10,0],[5,5]]"), false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, isConvex(tc.points))
		})
	}
}

func isConvex(points [][]int) bool {
	// Gift-wrapping approach
	// For each triplet of points (p1, p2, p3) in the polygon,
	// compute the z-component of the cross product of the vectors defined by the
	// edges pointing towards the points in increasing order.
	// Either all cross products are positive or negative, otherwise the polygon
	// isnon-convex.

	points = append(points, points[0], points[1])
	var allPos bool
	for i := 0; i < len(points)-2; i++ {
		p1, p2, p3 := points[i], points[i+1], points[i+2]
		dx1 := p2[0] - p1[0]
		dy1 := p2[1] - p1[1]
		dx2 := p3[0] - p2[0]
		dy2 := p3[1] - p2[1]
		zcross := dx1*dy2 - dy1*dx2
		if i == 0 {
			if zcross > 0 {
				allPos = true
			}
		}
		if allPos && zcross < 0 {
			return false
		} else if !allPos && zcross > 0 {
			return false
		}
	}
	return true
}
