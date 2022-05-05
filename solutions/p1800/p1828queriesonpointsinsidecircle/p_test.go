package p1828queriesonpointsinsidecircle

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPoints(t *testing.T) {
	for _, tc := range []struct {
		points  [][]int
		queries [][]int
		want    []int
	}{
		{[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}, []int{3, 2, 2}},
		{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}, []int{2, 3, 2, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, countPoints(tc.points, tc.queries))
		})
	}
}

func countPoints(points [][]int, queries [][]int) []int {
	// points = [x, y]
	// queries = [x, y, r]
	// sort points by x
	n := len(points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// For a point to fall within the queried circle, it must
	// at least fall within [x-r, x+r] and [y-r, y+r]
	res := make([]int, len(queries))
	for queryIdx, q := range queries {
		x, y, r := q[0], q[1], q[2]
		xStart := sort.Search(n, func(i int) bool {
			return points[i][0] >= x-r
		})
		for i := xStart; i < n && points[i][0] <= x+r; i++ {
			px, py := points[i][0], points[i][1]
			px -= x
			py -= y
			if px*px+py*py <= r*r {
				res[queryIdx]++
			}
		}
	}

	return res
}
