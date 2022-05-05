package p0587erectthefence

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_outerTrees(t *testing.T) {
	for _, tc := range []struct {
		trees [][]int
		want  [][]int
	}{
		{[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}, [][]int{{1, 1}, {2, 0}, {3, 3}, {2, 4}, {4, 2}}},
		{[][]int{{1, 2}, {2, 2}, {4, 2}}, [][]int{{4, 2}, {2, 2}, {1, 2}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.trees), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, outerTrees(tc.trees))
		})
	}
}

type Point struct {
	x, y int16
}

func (p Point) DistanceTo(p2 Point) int16 {
	dy := p2.y - p.y
	dx := p2.x - p.x
	return dy*dy + dx*dx
}

func outerTrees(trees [][]int) [][]int {
	points := make([]Point, len(trees))
	for i := range trees {
		points[i] = Point{int16(trees[i][0]), int16(trees[i][1])}
	}

	// Sort points by x-value (and in case of tie, lower y wins)
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	// Traverse points from left to right, adding to the hull
	hull := make([]Point, 0)
	n := 0
	for _, point := range points {
		for len(hull) >= 2 && orientation(hull[n-2], hull[n-1], point) > 0 {
			hull = hull[:n-1]
			n--
		}
		hull = append(hull, point)
		n++
	}
	hull = hull[:n-1]
	n--
	for i := len(points) - 1; i >= 0; i-- {
		point := points[i]
		for len(hull) >= 2 && orientation(hull[n-2], hull[n-1], point) > 0 {
			hull = hull[:n-1]
			n--
		}
		hull = append(hull, point)
		n++
	}

	// Remove redundant points
	res := make([][]int, 0, len(hull))
	seen := make(map[Point]struct{}, len(hull))
	for _, p := range hull {
		if _, exists := seen[p]; exists {
			continue
		}
		seen[p] = struct{}{}
		res = append(res, []int{int(p.x), int(p.y)})
	}

	return res
}

// See https://www.geeksforgeeks.org/orientation-3-ordered-points/
//
// Orientation returns a positive value if the curve formed by p1->p2->p3 is
// counter-clockwise, i.e. a left turn. If orientation returns zero, then the
// points are collinear (in a line). If orientation returns a positive value,
// then the three points form a clockwise (right) turn.
//
// The calculation can be derived like so:
//
// If the slope from p1 to p2 is smaller than that of p2 to p3, then the line
// a counter-clockwise turn. The slope from p1->p2 can be calculated by
// (p2_y - p1_y)/(p2_x - p1_x), and vice versa for p2->p3.
// If the slope p1->p2 is smaller than p2->p3, then p1p2 < p2p3,
// <=> p1p2 - p2p3 < 0
// The denominators can be factored out, resulting in the formula:
// (p2_y - p1_y)*(p3_x - p2_x) - (p3_y - p2_y)*(p2_x - p1_x) < 0
func orientation(p1, p2, p3 Point) int16 {
	return (p2.y-p1.y)*(p3.x-p2.x) - (p3.y-p2.y)*(p2.x-p1.x)
}
