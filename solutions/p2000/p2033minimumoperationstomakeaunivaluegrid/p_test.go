package p2033minimumoperationstomakeaunivaluegrid

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		x    int
		want int
	}{
		{[][]int{{980, 476, 644, 56}, {644, 140, 812, 308}, {812, 812, 896, 560}, {728, 476, 56, 812}}, 84, 45},
		{[][]int{{2, 4}, {6, 8}}, 2, 4},
		{[][]int{{1, 5}, {2, 3}}, 1, 5},
		{[][]int{{1, 2}, {3, 4}}, 2, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.grid, tc.x))
		})
	}
}

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	// Collect all values into a big list
	vals := make([]int, 0, m*n)
	for i := range grid {
		for j := range grid[i] {
			vals = append(vals, grid[i][j])
		}
	}
	// Pick any number from vals and check so that all other numbers are evenly
	// divisible by x
	for i := 1; i < len(vals); i++ {
		if abs(vals[i]-vals[0])%x != 0 {
			return -1
		}
	}

	// Sort list
	sort.Ints(vals)

	// Pick median
	median := vals[len(vals)/2]
	minCost := cost(vals, median, x)
	if len(vals)%2 == 0 {
		minCost = min(minCost, cost(vals, vals[(len(vals)/2)+1], x))
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cost(vals []int, targetVal, delta int) int {
	var cost int
	for _, val := range vals {
		cost += abs(val-targetVal) / delta
	}
	return cost
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
