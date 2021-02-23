package p0174dungeongame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculateMinimumHP(t *testing.T) {
	for _, tc := range []struct {
		dungeon [][]int
		want    int
	}{
		{[][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}, 3},
		// {[][]int{{2}, {1}}, 1},
		// {[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dungeon), func(t *testing.T) {
			require.Equal(t, tc.want, calculateMinimumHP(tc.dungeon))
		})
	}
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

func calculateMinimumHP(dungeon [][]int) int {
	// Idea: calculate the minimum hp required to reach the last square
	// Then continue to calculate the minimum hp to reach the previous square,
	// and so on, until we've reached the first square.
	m, n := len(dungeon), len(dungeon[0])
	requiredHP := make([][]int, m)
	for i := range requiredHP {
		requiredHP[i] = make([]int, n)
	}

	// The only option for the right and bottom edge is to continue
	// along the edge, this is the base-case
	requiredHP[m-1][n-1] = max(1, -dungeon[m-1][n-1]+1)
	for j := n - 2; j >= 0; j-- {
		requiredHP[m-1][j] = max(1, requiredHP[m-1][j+1]-dungeon[m-1][j])
	}
	for i := m - 2; i >= 0; i-- {
		requiredHP[i][n-1] = max(1, requiredHP[i+1][n-1]-dungeon[i][n-1])
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			requiredHP[i][j] = max(1, min(requiredHP[i][j+1], requiredHP[i+1][j])-dungeon[i][j])
		}
	}

	return requiredHP[0][0]
}
