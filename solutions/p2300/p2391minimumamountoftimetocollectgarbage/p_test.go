package p2391minimumamountoftimetocollectgarbage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_garbageCollection(t *testing.T) {
	for _, tc := range []struct {
		garbage []string
		travel  []int
		want    int
	}{
		{[]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}, 21},
		{[]string{"MMM", "PGM", "GP"}, []int{3, 10}, 37},
	} {
		t.Run(fmt.Sprintf("%+v", tc.garbage), func(t *testing.T) {
			require.Equal(t, tc.want, garbageCollection(tc.garbage, tc.travel))
		})
	}
}

func garbageCollection(garbage []string, travel []int) int {
	// The result will be determined by the shortest time for a truck to finish
	// all houses.
	// Let's just calculate it
	n := len(garbage)
	countUnits := func(s string, r rune) int {
		var res int
		for _, ch := range s {
			if ch == r {
				res++
			}
		}
		return res
	}
	// Time[0] = time without going
	var time [3][2]int
	for i, ch := range []rune{'G', 'P', 'M'} {
		// Skip any trailing houses
		j := n - 1
		for j >= 0 && countUnits(garbage[j], ch) == 0 {
			j--
		}
		// Count time
		for k := 0; k <= j; k++ {
			time[i][0] += countUnits(garbage[k], ch)
			if k < j {
				time[i][0] += travel[k]
			}
		}
	}
	return time[0][0] + time[1][0] + time[2][0]
}
