package p0435nonoverlappingintervals

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"

	"github.com/stretchr/testify/require"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      int
	}{
		{leetcode.ParseMatrix("[[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]]"), 7},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, eraseOverlapIntervals(tc.intervals))
		})
	}
}

func eraseOverlapIntervals(intervals [][]int) int {
	// Sort intervals by end-time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// Greedily pick the interval that ends first
	var removed int
	currentEnd := math.MinInt32
	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if start < currentEnd {
			removed++
			continue
		}
		currentEnd = end
	}
	return removed
}
