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
		{[][]int{{2}, {1}}, 1},
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.dungeon), func(t *testing.T) {
			require.Equal(t, tc.want, calculateMinimumHP(tc.dungeon))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	required := make([][]int, m)
	current := make([][]int, m)
	for i := range required {
		required[i] = make([]int, n)
		current[i] = make([]int, n)
		if i == 0 {
			if dungeon[0][0] < 0 {
				required[0][0] = -dungeon[0][0]
				current[0][0] = 0 // for clarity
			} else {
				required[0][0] = 0
				current[0][0] = dungeon[0][0]
			}
			continue
		}
		current[i][0] = current[i-1][0] + dungeon[i][0]
		if current[i][0] < 0 {
			required[i][0] = required[i-1][0] - current[i][0]
			current[i][0] = 0
		} else {
			required[i][0] = required[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		current[0][j] = current[0][j-1] + dungeon[0][j]
		if current[0][j] < 0 {
			required[0][j] = required[0][j-1] - current[0][j]
			current[0][j] = 0
		} else {
			required[0][j] = required[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := dungeon[i][j]
			// left
			var leftReq, leftHealth int
			leftDiff := v + current[i][j-1]
			if leftDiff < 0 {
				leftHealth = 0
				leftReq = required[i][j-1] - leftDiff
			} else {
				leftHealth = v + current[i][j-1]
				leftReq = required[i][j-1]
			}

			var aboveReq, aboveHealth int
			aboveDiff := v + current[i-1][j]
			if aboveDiff < 0 {
				// Have to add to required to come from above
				aboveHealth = 0
				aboveReq = required[i-1][j] - aboveDiff
			} else {
				aboveHealth = v + current[i-1][j]
				aboveReq = required[i-1][j]
			}

			if aboveReq < leftReq {
				required[i][j] = aboveReq
				current[i][j] = aboveHealth
			} else if leftReq < aboveReq {
				required[i][j] = leftReq
				current[i][j] = leftHealth
			} else {
				required[i][j] = leftReq // arbitrary choice, both are equal
				if leftHealth > aboveHealth {
					current[i][j] = leftHealth
				} else {
					current[i][j] = aboveHealth
				}
			}
		}
	}

	return 1 + required[m-1][n-1]
}
