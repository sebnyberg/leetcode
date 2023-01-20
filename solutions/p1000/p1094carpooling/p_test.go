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
	type delta struct {
		t int
		d int
	}
	deltas := []delta{}
	for _, t := range trips {
		deltas = append(deltas, delta{t: t[1], d: t[0]})
		deltas = append(deltas, delta{t: t[2], d: -t[0]})
	}
	sort.Slice(deltas, func(i, j int) bool {
		if deltas[i].t == deltas[j].t {
			return deltas[i].d < 0
		}
		return deltas[i].t < deltas[j].t
	})

	for _, d := range deltas {
		capacity -= d.d
		if capacity < 0 {
			return false
		}
	}
	return true
}
