package p0120triangle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTotal(t *testing.T) {
	for _, tc := range []struct {
		triangle [][]int
		want     int
	}{
		{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
	} {
		t.Run(fmt.Sprintf("%+v", tc.triangle), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTotal(tc.triangle))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) (res int) {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	cost := triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			cost[j] = min(cost[j], cost[j+1]) + triangle[i][j]
		}
	}

	return cost[0]
}
