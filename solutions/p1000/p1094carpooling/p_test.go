package p1094carpooling

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_carPooling(t *testing.T) {
	for _, tc := range []struct {
		trips    [][]int
		capacity int
		want     bool
	}{
		{[][]int{{4, 1, 3}, {1, 3, 6}, {8, 2, 3}, {8, 4, 6}, {4, 4, 8}}, 12, false},
		{[][]int{{5, 1, 3}, {5, 3, 6}, {1, 1, 2}, {9, 2, 5}, {6, 5, 7}, {2, 3, 6}, {9, 3, 5}}, 26, true},
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 4, false},
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 5, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.trips), func(t *testing.T) {
			require.Equal(t, tc.want, carPooling(tc.trips, tc.capacity))
		})
	}
}

func carPooling(trips [][]int, capacity int) bool {
	type change struct {
		delta int16
		time  int16
	}

	changes := make([]change, 0, len(trips)*2)
	for _, t := range trips {
		changes = append(changes,
			change{int16(-t[0]), int16(t[1])},
			change{int16(t[0]), int16(t[2])},
		)
	}
	sort.Slice(changes, func(i, j int) bool {
		return changes[i].time < changes[j].time
	})

	for i, ch := range changes {
		capacity += int(ch.delta)
		if i < len(changes)-1 && changes[i].time != changes[i+1].time && capacity < 0 {
			return false
		}
	}
	return true
}
