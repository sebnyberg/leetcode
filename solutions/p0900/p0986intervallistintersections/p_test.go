package p0986intervallistintersections

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intervalIntersection(t *testing.T) {
	for _, tc := range []struct {
		firstList  [][]int
		secondList [][]int
		want       [][]int
	}{
		{
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.firstList), func(t *testing.T) {
			require.Equal(t, tc.want, intervalIntersection(tc.firstList, tc.secondList))
		})
	}
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// Merge the two into one list of +1 and -1 first, and add intervals
	// whenever the two overlap
	deltas := make([][2]int, 0)
	for _, x := range firstList {
		deltas = append(deltas, [2]int{x[0], 1}, [2]int{x[1], -1})
	}
	for _, x := range secondList {
		deltas = append(deltas, [2]int{x[0], 2}, [2]int{x[1], -2})
	}
	sort.Slice(deltas, func(i, j int) bool {
		a := deltas[i]
		b := deltas[j]
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return a[0] < b[0]
	})
	var l int
	var cur int
	var res [][]int
	for _, d := range deltas {
		if cur == 3 {
			res = append(res, []int{l, d[0]})
		}
		cur ^= abs(d[1])
		if cur == 3 {
			l = d[0]
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
