package p2332thelatesttimetocatchabus

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_latestTimeCatchTheBus(t *testing.T) {
	for _, tc := range []struct {
		buses      []int
		passengers []int
		capacity   int
		want       int
	}{
		{
			[]int{18, 8, 3, 12, 9, 2, 7, 13, 20, 5},
			[]int{13, 10, 8, 4, 12, 14, 18, 19, 5, 2, 30, 34},
			1, 11,
		},
		{[]int{3}, []int{2, 4}, 2, 3},
		{[]int{3}, []int{2, 4}, 2, 3},
		{[]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2, 20},
		{[]int{2}, []int{2}, 1, 1},
		{[]int{10, 20}, []int{2, 17, 18, 19}, 2, 16},
	} {
		t.Run(fmt.Sprintf("%+v", tc.buses), func(t *testing.T) {
			require.Equal(t, tc.want, latestTimeCatchTheBus(tc.buses, tc.passengers, tc.capacity))
		})
	}
}

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	// I found this to be crazy hard in terms of logical reasoning.
	// The idea is to try to fill buses one by one.
	// Then for each bus, check which passengers could go on that bus.
	// Then determine the best time where we cound enter the bus (if any).
	// Typically, we try to squeeze ourselves into the bus right before the last
	// passenger. If many passengers enter in sequence, then this may not be
	// possible.
	sort.Ints(buses)
	sort.Ints(passengers)

	var k int  // first passenger on the bus (if any)
	var kk int // first passenger not on the bus

	// For each bus
	var res int
	var i int
	for kk = k; i < len(buses); k, i = kk, i+1 {
		// Check which passengers could be added
		for kk < len(passengers) && passengers[kk] <= buses[i] &&
			kk-k < capacity {
			kk++
		}
		count := kk - k

		// Check whether we can enter right as the bus leaves
		if count == 0 || count < capacity && passengers[kk-1] != buses[i] {
			res = buses[i]
			continue
		}

		// Find the latest time where we could enter the bus
		// If passengers are entering in sequence, there may be no such time.
		for m := k; m < kk; m++ {
			if m == 0 { // Can always enter before first passenger
				res = passengers[m] - 1
				continue
			}
			if passengers[m-1] != passengers[m]-1 {
				res = passengers[m] - 1
			}
		}
	}

	return res
}
