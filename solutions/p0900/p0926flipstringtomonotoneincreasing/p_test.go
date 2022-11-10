package p0926flipstringtomonotoneincreasing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minFlipsMonoIncr(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"00110", 1},
		{"010110", 2},
		{"00011000", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minFlipsMonoIncr(tc.s))
		})
	}
}

func minFlipsMonoIncr(s string) int {
	n := len(s)
	left := make([]int, n+1)
	for i := range s {
		left[i+1] = left[i]
		if s[i] == '1' {
			left[i+1]++
		}
	}
	minCost := n
	for j := 0; j <= len(s); j++ {
		// making everything on the left side into zeroes will cost
		x := left[j]
		// making everything on  the right into ones will cost
		y := n - j - (left[n] - left[j])
		cost := x + y
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}
