package p1883minskipstoarriveatmeetingontime

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSkips(t *testing.T) {
	for _, tc := range []struct {
		dist        []int
		speed       int
		hoursBefore int
		want        int
	}{
		{[]int{1, 3, 2}, 4, 2, 1},
		{[]int{7, 3, 5, 5}, 2, 10, 2},
		{[]int{7, 3, 5, 5}, 1, 10, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dist), func(t *testing.T) {
			require.Equal(t, tc.want, minSkips(tc.dist, tc.speed, tc.hoursBefore))
		})
	}
}

const eps = 1e-9

func minSkips(dist []int, speed int, hoursBefore int) int {
	// For each distance pair, there is always
	// There is up to 1000 distances
	// dp[i][j] = min distance at road i having done j skips
	var dp [1001][1001]float64
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = math.MaxFloat64
		}
	}
	dp[0][0] = 0

	n := len(dist)
	for i := 1; i <= n; i++ {
		// Don't use speedup
		dp[i][0] = math.Ceil(dp[i-1][0] + float64(dist[i-1])/float64(speed) - eps)
		for j := 1; j < i+1; j++ {
			// Either use speedup, or
			withSpeedup := dp[i-1][j-1] + float64(dist[i-1])/float64(speed)
			// Don't use speedup previous result + current
			previousSpeedup := math.Ceil(dp[i-1][j] + float64(dist[i-1])/float64(speed) - eps)
			dp[i][j] = min(withSpeedup, previousSpeedup)
		}
	}
	for j := range dp[n] {
		if dp[n][j] <= float64(hoursBefore) {
			return j
		}
	}
	return -1
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
