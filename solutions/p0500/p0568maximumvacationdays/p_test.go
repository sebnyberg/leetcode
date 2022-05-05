package p0568maximumvacationdays

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxVacationDays(t *testing.T) {
	for _, tc := range []struct {
		flights [][]int
		days    [][]int
		want    int
	}{
		{[][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, [][]int{{1, 3, 1}, {6, 0, 3}, {3, 3, 3}}, 12},
		{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, [][]int{{1, 1, 1}, {7, 7, 7}, {7, 7, 7}}, 3},
		{[][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, [][]int{{7, 0, 0}, {0, 7, 0}, {0, 0, 7}}, 21},
	} {
		t.Run(fmt.Sprintf("%+v", tc.flights), func(t *testing.T) {
			require.Equal(t, tc.want, maxVacationDays(tc.flights, tc.days))
		})
	}
}

func maxVacationDays(flights [][]int, days [][]int) int {
	// You are in city indexed 0 on Monday
	// Flights matrix denotes whether it's possible to fly from i to j, in which
	// case the corresponding value is 1 (otherwise 0)
	// Can only take flights on each week's Monday morning

	// There appears to be a risk of getting stuck in a city

	// It does not matter how flights happend prior to a given day, only what the
	// total number of vacation days are.

	// This tells us that this is a DP-exercise, and the maximum number of days
	// in any position is the max days for any place that can reach that place
	// for the prior day.

	n := len(flights)
	reachableFrom := make([][]int, n)
	for from := range flights {
		// Can stay in the same city
		reachableFrom[from] = append(reachableFrom[from], from)

		// And be reached by flights
		for to := range flights[from] {
			if flights[from][to] == 1 {
				reachableFrom[to] = append(reachableFrom[to], from)
			}
		}
	}

	prev := make([]int, n)
	for i := range prev {
		prev[i] = math.MinInt32
	}
	prev[0] = 0
	cur := make([]int, n)
	for t := 0; t < len(days[0]); t++ {
		// Zero-out current
		for i := range cur {
			cur[i] = math.MinInt32
		}
		for f := range flights {
			for _, from := range reachableFrom[f] {
				cur[f] = max(cur[f], prev[from]+days[f][t])
			}
		}
		prev, cur = cur, prev
	}
	var maxVal int
	for _, v := range prev {
		maxVal = max(maxVal, v)
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
