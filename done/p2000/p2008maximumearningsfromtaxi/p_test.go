package p2008maximumearningfromtaxi

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxTaxiEarnings(t *testing.T) {
	for _, tc := range []struct {
		n     int
		rides [][]int
		want  int64
	}{
		{5, [][]int{{2, 5, 4}, {1, 5, 1}}, 7},
		{20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}, 20},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxTaxiEarnings(tc.n, tc.rides))
		})
	}
}

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] < rides[j][0]
	})
	dp := make([]int, n+1)
	var rideIdx int
	for i := 1; i <= n; i++ {
		dp[i] = max(dp[i], dp[i-1])
		for rideIdx < len(rides) && rides[rideIdx][0] == i {
			r := rides[rideIdx]
			start, end, tip := r[0], r[1], r[2]
			dp[end] = max(dp[end], dp[start]+end-start+tip)
			rideIdx++
		}
	}
	return int64(dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
