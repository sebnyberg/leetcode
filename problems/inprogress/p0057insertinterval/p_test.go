package p0057insertinterval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_insert(t *testing.T) {
	for _, tc := range []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{{1, 5}}, []int{0, 3}, [][]int{{0, 5}}},
		{[][]int{{1, 5}, {10, 11}}, []int{6, 7}, [][]int{{1, 5}, {6, 7}, {10, 11}}},
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
		{[][]int{{1, 5}}, []int{2, 7}, [][]int{{1, 7}}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.intervals, tc.newInterval), func(t *testing.T) {
			require.Equal(t, tc.want, insert(tc.intervals, tc.newInterval))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	res := make([][]int, 0, n)
	if newInterval[1] < intervals[0][0] {
		intervals = append(intervals, []int{})
		copy(intervals[1:], intervals)
		intervals[0] = newInterval
		return intervals
	}
	if newInterval[0] > intervals[n-1][1] {
		return append(intervals, newInterval)
	}

	var i int
	for ; i < len(intervals) && newInterval[0] > intervals[i][1]; i++ {
		res = append(res, intervals[i])
	}

	if newInterval[1] < intervals[i][0] {
		res = append(res, newInterval)
		return append(res, intervals[i:]...)
	}

	start, end := min(newInterval[0], intervals[i][0]), max(intervals[i][1], newInterval[1])
	for i++; i < len(intervals) && intervals[i][0] <= end; i++ {
		end = max(end, intervals[i][1])
	}
	res = append(res, []int{start, end})
	if i < len(intervals) {
		res = append(res, intervals[i:]...)
	}
	return res
}
