package p2188mintimetofinishtherace

import (
	"fmt"
	"leetcode"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumFinishTime(t *testing.T) {
	for _, tc := range []struct {
		tires      [][]int
		changeTime int
		numLaps    int
		want       int
	}{
		// {leetcode.ParseMatrix("[[1,2],[1,2],[1,2]]"), 5, 4, 21},
		{leetcode.ParseMatrix("[[2,3],[3,4]]"), 5, 4, 21},
		{leetcode.ParseMatrix("[[1,10],[2,2],[3,4]]"), 6, 5, 25},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tires), func(t *testing.T) {
			require.Equal(t, tc.want, minimumFinishTime(tc.tires, tc.changeTime, tc.numLaps))
		})
	}
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	// Some things are for sure - changing a tire for which the changeTime +
	// current lap time for the tire is greater than the original time is always
	// optimal.

	// When the current lap time is smaller than changeTime, then it's always
	// optimal to keep going (?)

	// Consider each tire one by one. Find the optimal change rate per tire, then
	// use it to combine the different tires into an optimal race.

	// I'll assume that it's only worth it to change tires if changeTime + origTime
	// is smaller than current time. I'm not sure this holds for all cases..

	// First, remove duplicate tires...
	sort.Slice(tires, func(i, j int) bool {
		if tires[i][0] == tires[j][0] {
			return tires[i][1] <= tires[j][1]
		}
		return tires[i][0] < tires[j][0]
	})
	var j int
	for i := 1; i < len(tires); i++ {
		if tires[i][0] == tires[j][0] && tires[i][1] == tires[j][1] {
			continue
		}
		j++
		tires[j] = tires[i]
	}
	tires = tires[:j+1]

	n := len(tires)
	presum := make([][]int, n)
	for i := range presum {
		presum[i] = make([]int, numLaps+1)
	}
	lapTimes := make([]int, 0, numLaps)
	for i, tire := range tires {
		lapTimes = lapTimes[:0]
		f, r := tire[0], tire[1]
		rr := 1
		for t := 0; t < numLaps; t++ {
			tot := f * rr
			if changeTime+f <= tot { // change tire
				rr = 1
				tot = changeTime + f
			}
			lapTimes = append(lapTimes, tot)
			rr *= r
		}
		for j := range lapTimes {
			presum[i][j+1] = presum[i][j] + lapTimes[j]
		}
	}

	// Now presum[i] contains the cost of keeping tire i throughout the entire
	// race. presum is guaranteed to cover all laps.

	// dp[i] = minimum cost to go i laps
	dp := make([]int, numLaps)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < len(dp); i++ {
		// For each lap i, the baseline cost is the minimum cost to keep the same
		// tire (any tire) throughout laps 0 to i
		for j := range presum {
			dp[i] = min(dp[i], presum[j][i+1])
		}

		// Otherwise, the current tire may be combined with some other tire to form
		// a more optimal solution
		for k := 0; k < i; k++ {
			before := dp[k]
			after := dp[i-k-1]
			dp[i] = min(dp[i], before+after+changeTime)
		}
	}

	// For each time during the race, the minimum possible time is equal to the
	return dp[numLaps-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
