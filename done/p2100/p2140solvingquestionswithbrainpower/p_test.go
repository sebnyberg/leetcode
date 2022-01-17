package p2140solvingquestionswithbrainpower

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mostPoints(t *testing.T) {
	for _, tc := range []struct {
		questions [][]int
		want      int64
	}{
		{[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, 5},
		{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.questions), func(t *testing.T) {
			require.Equal(t, tc.want, mostPoints(tc.questions))
		})
	}
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		points, brainPower := questions[i][0], questions[i][1]
		dp[i] = max(dp[i+1], points+dp[min(n, i+brainPower+1)])
	}
	return int64(dp[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
