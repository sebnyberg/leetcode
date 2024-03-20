package p0057insertinterval

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_insert(t *testing.T) {
	for _, tc := range []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{[][]int{{1, 2}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {4, 8}, {12, 16}}},
		{[][]int{{1, 2}, {4, 8}}, []int{12, 16}, [][]int{{1, 2}, {4, 8}, {12, 16}}},
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

func insert(intervals [][]int, newInterval []int) [][]int {
	i := sort.Search(len(intervals), func(j int) bool {
		return intervals[j][1] >= newInterval[0]
	})
	if i == len(intervals) {
		intervals = append(intervals, newInterval)
		return intervals
	}
	// i now points to the lowest interval that has a end >= newInterval.
	//
	// Case 1. The current interval is outside the newInterval
	if intervals[i][0] > newInterval[1] {
		// Then, we insert newInterval where intervals[i] was and return
		intervals = append(intervals, []int{})
		copy(intervals[i+1:], intervals[i:])
		intervals[i] = newInterval
		return intervals
	}
	left := i
	start := min(newInterval[0], intervals[i][0])
	end := max(newInterval[1], intervals[i][1])
	i++
	for i < len(intervals) && end >= intervals[i][0] {
		end = max(end, intervals[i][1])
		i++
	}
	// Replace the range [left:i] with the interval []int{start, end}
	m := i - left
	copy(intervals[left+1:], intervals[i:])
	intervals = intervals[:len(intervals)-(m-1)]
	intervals[left] = []int{start, end}
	return intervals
}
