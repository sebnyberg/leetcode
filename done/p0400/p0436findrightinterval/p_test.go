package p0435nonoverlappingintervals

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRightInterval(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      []int
	}{
		{[][]int{{1, 2}}, []int{-1}},
		{[][]int{{3, 4}, {2, 3}, {1, 2}}, []int{-1, 0, 1}},
		{[][]int{{1, 4}, {2, 3}, {3, 4}}, []int{-1, 2, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, findRightInterval(tc.intervals))
		})
	}
}

type intervalSlice struct {
	items [][]int
	idx   []int
}

func (s *intervalSlice) Less(i, j int) bool {
	return s.items[i][0] < s.items[j][0]
}

func (s *intervalSlice) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func (s *intervalSlice) Len() int { return len(s.idx) }

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)

	// Sort intervals by start
	sortedIntervals := &intervalSlice{
		items: intervals,
		idx:   make([]int, n),
	}
	for i := range intervals {
		sortedIntervals.idx[i] = i
	}
	sort.Sort(sortedIntervals)

	// For each interval, find the smallest start >= end of the current interval
	// using binary search.
	res := make([]int, n)
	for i, interval := range sortedIntervals.items {
		pos := sort.Search(n, func(i int) bool {
			return sortedIntervals.items[i][0] >= interval[1]
		})
		i := sortedIntervals.idx[i]
		if pos == n {
			res[i] = -1
		} else {
			res[i] = sortedIntervals.idx[pos]
		}
	}
	return res
}
