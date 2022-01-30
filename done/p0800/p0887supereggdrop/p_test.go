package p0887supereggdrop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_superEggDrop(t *testing.T) {
	for _, tc := range []struct {
		k    int
		n    int
		want int
	}{
		{1, 3, 3},
		{1, 2, 2},
		{2, 6, 3},
		{3, 14, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, superEggDrop(tc.k, tc.n))
		})
	}
}

func superEggDrop(k int, n int) int {
	// With k = 1, the only way to find out which floor the egg can be dropped
	// from is to try every floor from 1 until it breaks.
	// With k = 2, we may potentially eliminate half of the floors by dropping
	// the egg from the middle floor. If it breaks, then we have to try
	// floor-by-floor from 1 and up. If it does not break, then try from middle+1
	// and up.
	// Essentially, we can eliminate half of the floors per egg drop.
	//
	// If we drop an egg, there are two cases:
	//
	// 1. The egg does not break, we are left with the same amount of eggs
	// 2. The egg breaks, we lose one egg
	//
	// In case (1), we get superEggDrop(k, n-x) where x is the floor
	// In case (2), we get superEggDrop(k-1, x-1)
	//
	// In both cases, we have used one move.
	//
	// Let's explore the maximum amount of floors that are possible to reach
	// given m moves and k eggs.
	//
	// Given 1 move, it is only possible to explore one floor.
	//
	// maxFloor[k] = max floor reachable with k eggs at current number of moves
	maxFloor := make([]int, k+1)
	m := 0
	for maxFloor[k] < n {
		m++
		// For any move past 1 move, the max number of possible floors is equal to
		// 1 (the floor where the egg is dropped) + max moves given that the egg
		// breaks + max moves given that the egg does not break.
		//
		// This gives the recursion
		// maxFloor[k] = maxFloor[k-1]+maxFloor[k]+1
		//
		// Since this would override the existing element, we iterate from end to
		// start.
		for i := k; i > 0; i-- {
			maxFloor[i] = maxFloor[i-1] + maxFloor[i] + 1
		}
	}

	return m
}
