package p2087mincosthomecomingofarobotinagrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for _, tc := range []struct {
		startPos []int
		homePos  []int
		rowCosts []int
		colCosts []int
		want     int
	}{
		{[]int{1, 0}, []int{2, 3}, []int{5, 4, 3}, []int{8, 2, 6, 7}, 18},
		{[]int{0, 0}, []int{0, 0}, []int{5}, []int{26}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startPos), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.startPos, tc.homePos, tc.rowCosts, tc.colCosts))
		})
	}
}

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	// It may not be intuitive at first glance but the only possible variations
	// in solutions is which way to enter the house from (four alternatives)
	// To make the problem simpler to handle, we transpose costs and positions
	// so that the house is always to the bottom-right of the robot.
	m := len(rowCosts)
	n := len(colCosts)
	startX, startY := startPos[1], startPos[0]
	homeX, homeY := homePos[1], homePos[0]
	if startX > homeX {
		// Transpose x
		startX = n - startX - 1
		homeX = n - homeX - 1
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			colCosts[l], colCosts[r] = colCosts[r], colCosts[l]
		}
	}
	if startY > homeY {
		// Transpose y
		startY = m - startY - 1
		homeY = m - homeY - 1
		for l, r := 0, m-1; l < r; l, r = l+1, r-1 {
			rowCosts[l], rowCosts[r] = rowCosts[r], rowCosts[l]
		}
	}
	var cost int
	// Walk in y-direction until on house row
	for i := startY; i < homeY; i++ {
		cost += rowCosts[i+1]
	}
	// Walk in the x-direction until on house col
	for i := startX; i < homeX; i++ {
		cost += colCosts[i+1]
	}
	return cost
}
