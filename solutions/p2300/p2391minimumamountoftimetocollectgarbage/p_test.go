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
	// This problem was a bit confusing in its description. There is no "minimal"
	// time - it's just the total time for all trucks to do their work.
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
	var time int
	for _, ch := range []rune{'G', 'P', 'M'} {
		// Skip any trailing houses
		j := n - 1
		for j >= 0 && countUnits(garbage[j], ch) == 0 {
			j--
		}
		// Count time
		for k := 0; k <= j; k++ {
			time += countUnits(garbage[k], ch)
			if k < j {
				time += travel[k]
			}
		}
	}
	return time
}
