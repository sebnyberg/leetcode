package p0056mergeintervals

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, merge(tc.intervals))
		})
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	start, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] <= end {
			end = max(end, interval[1])
		} else {
			res = append(res, []int{start, end})
			start, end = interval[0], interval[1]
		}
	}
	return append(res, []int{start, end})
}
