package p2187mintimetocompletetrips

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTime(t *testing.T) {
	for _, tc := range []struct {
		time       []int
		totalTrips int
		want       int64
	}{
		{[]int{39, 82, 69, 37, 78, 14, 93, 36, 66, 61, 13, 58, 57, 12, 70, 14, 67, 75, 91, 1, 34, 68, 73, 50, 13, 40, 81, 21, 79, 12, 35, 18, 71, 43, 5, 50, 37, 16, 15, 6, 61, 7, 87, 43, 27, 62, 95, 45, 82, 100, 15, 74, 33, 95, 38, 88, 91, 47, 22, 82, 51, 19, 10, 24, 87, 38, 5, 91, 10, 36, 56, 86, 48, 92, 10, 26, 63, 2, 50, 88, 9, 83, 20, 42, 59, 55, 8, 15, 48, 25}, 4187, 858},
		{[]int{5, 10, 10}, 9, 25},
		{[]int{1, 2, 3}, 5, 3},
		{[]int{2}, 1, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.time), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTime(tc.time, tc.totalTrips))
		})
	}
}

func minimumTime(time []int, totalTrips int) int64 {
	// Use binary search to find the first time for which the total trips is
	// totalTrips
	var lo, hi = 0, math.MaxInt64
	for lo < hi {
		mid := (lo + hi) / 2
		var trips int
		for _, t := range time {
			if trips > totalTrips {
				break
			}
			tripsForBus := mid / t
			trips += tripsForBus
		}
		if trips >= totalTrips {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return int64(lo)
}
