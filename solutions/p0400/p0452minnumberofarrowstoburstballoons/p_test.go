package p0452minnumberofarrowstoburstballoons

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMinArrowShots(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, findMinArrowShots(tc.points))
		})
	}
}

func findMinArrowShots(points [][]int) int {
	// We need to pop every balloon in the list of balloons
	// Thus, there must exist a point x between the two edges
	// of each interval that is the most efficiant place to shoot for.
	//
	// Idea: sort by x_start
	// * Sort by x_start
	// * While there are new items in points for which x_start is
	// 	less than or equal to the smallest x_end so far, keep poppin'
	// Repeat process to get the solution
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	n := len(points)
	var i int
	var arrows int
	for i < n {
		end := points[i][1]
		i++
		for i < n && points[i][0] <= end {
			end = min(end, points[i][1])
			i++
		}
		arrows++
	}
	return arrows
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
