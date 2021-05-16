package p1864minnumberofswapstomakebinstringalternating

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwaps(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"010", 0},
		{"111000", 1},
		{"1110", -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minSwaps(tc.s))
		})
	}
}

func minSwaps(s string) int {
	// If the difference between ones and zeroe is greater than one,
	// there is no solution
	var ones, zeroes int
	for _, ch := range s {
		if ch == '1' {
			ones++
		} else {
			zeroes++
		}
	}
	d := abs(ones - zeroes)
	if d > 1 {
		return -1
	}

	// Do a frequency count for even / odd positions
	var count [2][2]int
	for i, ch := range s {
		if i%2 == 0 {
			count[0][ch-'0']++
		} else {
			count[1][ch-'0']++
		}
	}

	// If there is a difference of 1, then the most common number
	// must be at the edge.
	if d == 1 {
		if ones > zeroes {
			// Ones go in even slots, ans is number of zeroes in even slots
			return count[0][0]
		} else {
			// Zeroes go in even slots, ans is number of ones in even slots
			return count[0][1]
		}
	}

	// Make min possible number of swaps
	return min(
		count[0][0], count[0][1],
	)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
