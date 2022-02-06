package p2167minimumtimetoremoveallcarscontainingillegalgoods

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTime(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"1100101", 5},
		{"0010", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTime(tc.s))
		})
	}
}

func minimumTime(s string) int {
	costLeft := make([]int, len(s)+1)
	costLeft[0] = 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '1' {
			costLeft[i+1] = min(costLeft[i]+2, i+1)
		} else {
			costLeft[i+1] = costLeft[i]
		}
	}
	costRight := make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == '1' {
			costRight[i] = min(costRight[i+1]+2, len(s)-i)
		} else {
			costRight[i] = costRight[i+1]
		}
	}
	minCost := math.MaxInt32
	for i := 0; i < len(s); i++ {
		minCost = min(minCost, costLeft[i]+costRight[i])
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
