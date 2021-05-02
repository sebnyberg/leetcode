package p1851minintervaltoincludeeachquery

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minInterval(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		queries   []int
		want      []int
	}{
		{[][]int{{6, 6}, {5, 5}, {10, 10}, {3, 6}, {9, 9}}, []int{4}, []int{4}},
		{[][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}, []int{3, 3, 1, 4}},
		{[][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}, []int{2, -1, 4, 6}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, minInterval(tc.intervals, tc.queries))
		})
	}
}

func minInterval(intervals [][]int, queries []int) []int {
	// Let's try something dumb
	// Sort intervals by interval's left edge
	// For each query, do binary search to find the first left edge that is
	// lower than the query
	// While the size of the smallest interval is greater than the
	// distance between the current left edge and the query position,
	// keep looking for new intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ni := len(intervals)
	nq := len(queries)
	res := make([]int, nq)
	for i, q := range queries {
		cand := sort.Search(ni, func(j int) bool {
			return intervals[j][0] > q
		})
		res[i] = -1
		cand--
		if cand < 0 {
			continue
		}
		minDist := math.MaxInt32
		for cand >= 0 && q-intervals[cand][0] <= minDist {
			if intervals[cand][1] >= q {
				minDist = min(minDist, intervals[cand][1]-intervals[cand][0]+1)
			}
			cand--
		}
		if minDist != math.MaxInt32 {
			res[i] = minDist
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
