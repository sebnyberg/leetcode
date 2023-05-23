package p1453maximumnumberofdartsinsideofacirculardartboard

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numPoints(t *testing.T) {
	for i, tc := range []struct {
		darts [][]int
		r     int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[4,-4],[-2,0],[0,2],[-3,1],[2,3],[2,4],[1,1]]"),
			3, 6,
		},
		{
			leetcode.ParseMatrix("[[-3,0],[3,0],[2,6],[5,4],[0,9],[7,8]]"),
			5, 5,
		},
		{
			leetcode.ParseMatrix("[[-2,0],[2,0],[0,2],[0,-2]]"),
			2, 4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numPoints(tc.darts, tc.r))
		})
	}
}

const eps = 1e-5

func numPoints(darts [][]int, r int) int {
	// For any circle, the following is true:
	//
	// 1. Two darts are on the border of the circle,
	// 2. Two darts could be on the border of the circle and still contain as
	// many darts.
	//
	// In other words, we only need to check pairs of points and count number of
	// darts inside the corresponding circle.
	//
	// The location for the centerpoint is given by the third point of the
	// triangle where two sides is of length "r" and the base is the distance
	// between the two points.
	//
	res := 1
	dist := func(a, b [2]float64) float64 {
		dx := b[0] - a[0]
		dy := b[1] - a[1]
		return math.Sqrt(dx*dx + dy*dy)
	}

	for i := 0; i < len(darts)-1; i++ {
		for j := i + 1; j < len(darts); j++ {
			// Given the two points A and B
			// There exists two possible third points C
			// Such that AC == BC
			ax := float64(darts[i][0])
			ay := float64(darts[i][1])
			bx := float64(darts[j][0])
			by := float64(darts[j][1])
			if ax == bx && ay == by {
				continue
			}
			if ax > bx {
				ax, bx = bx, ax
				ay, by = by, ay
			}
			dx := bx - ax
			dy := by - ay

			ab := math.Sqrt(dx*dx + dy*dy)
			if ab > 2*float64(r)+eps {
				continue
			}

			midx := (ax + bx) / 2
			midy := (ay + by) / 2
			midToCenter := math.Sqrt(float64(r*r) - (ab/2)*(ab/2))

			cx1 := midx - midToCenter*dy/ab
			cx2 := midx + midToCenter*dy/ab
			cy1 := midy + midToCenter*dx/ab
			cy2 := midy - midToCenter*dx/ab

			for _, p := range [][2]float64{{cx1, cy1}, {cx2, cy2}} {
				var count int
				for _, q := range darts {
					d := dist(p, [2]float64{float64(q[0]), float64(q[1])})
					if d-float64(r) <= eps {
						count++
					}
				}
				res = max(res, count)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
