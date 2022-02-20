package p1288removecoveredintervals

import (
	"fmt"
	"leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeCoveredIntervals(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      int
	}{
		{leetcode.ParseMatrix("[[34335,39239],[15875,91969],[29673,66453],[53548,69161],[40618,93111]]"), 2},
		{[][]int{{1, 4}, {3, 6}, {2, 8}}, 2},
		{[][]int{{1, 4}, {2, 3}}, 1},
		{[][]int{{3, 10}, {4, 10}, {5, 11}}, 2},
		{[][]int{{1, 2}, {1, 4}, {3, 4}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, removeCoveredIntervals(tc.intervals))
		})
	}
}

func removeCoveredIntervals(intervals [][]int) int {
	// If intervals is sorted by start, then all intervals at index j > i must
	// have a start-date that is "covered" by all previous interval start-dates.
	// The question is whether the end-date of any previous intervals end after
	// the current interval.
	is := intervals
	sort.Slice(is, func(i, j int) bool {
		if is[i][0] == is[j][0] {
			return is[i][1] > is[j][0]
		}
		return is[i][0] < is[j][0]
	})

	var cur int
	remaining := len(is)
	for i := 1; i < len(is); i++ {
		if is[i][1] <= is[cur][1] {
			remaining--
		}
		if is[i][1] > is[cur][1] {
			cur = i
		}
	}

	return remaining
}
