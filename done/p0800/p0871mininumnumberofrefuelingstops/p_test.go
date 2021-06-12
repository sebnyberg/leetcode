package p0871mininumnumberofrefuelingstops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minRefuelStops(t *testing.T) {
	for _, tc := range []struct {
		target    int
		startFuel int
		stations  [][]int
		want      int
	}{
		{1, 1, nil, 0},
		{100, 1, [][]int{{10, 100}}, -1},
		{100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, minRefuelStops(tc.target, tc.startFuel, tc.stations))
		})
	}
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	/*
		Intuition:
		* For each position in the race, it is unknown whether we should refill or not,
			giving a total of 2^500 possible refill/no-refill cases.
		* Since there are refills with up to 10^9, it is not possible to allocate an
			array containing the state at each possible position. I.e. only the points
			(start, stations[0], ,..., stations[n-1], end) can be considered
		* Start and end positions can be considered to be stations to simplify code.
		* It's not enough to consider the minimum number of stops between two gas
			stations - the amount of fuel left in the tank also matters.
		* If we can get to a station with more fuel in the tank and the same amount
			of stops, it is objectively better than a previous result with less fuel.
	*/
	stations = append(stations, []int{target, 0})
	n := len(stations)

	// Note: DP could be reduced to a one-dimensional slice, but I'll leave this
	// unoptimized version here.
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = startFuel // zero stops = startFuel
	for i, station := range stations {
		var dist int
		if i > 0 {
			dist = station[0] - stations[i-1][0]
		} else {
			dist = station[0]
		}
		// for each possible number of refills
		for j := 0; j <= i+1; j++ {
			dp[i+1][j] = dp[i][j] - dist // no refill here, fuel just goes down

			// if it is possible to reach this post with one less refill,
			// we may want to refill here
			if j > 0 && dp[i][j-1]-dist >= 0 {
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-1]+station[1]-dist)
			}
		}
	}

	// The first positive entry (if any) is the answer
	for stops, remainingGas := range dp[n] {
		if remainingGas >= 0 {
			return stops
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
