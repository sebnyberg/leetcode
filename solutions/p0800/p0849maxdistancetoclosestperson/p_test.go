package p0849maxdistancetoclosestperson

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDistToClosest(t *testing.T) {
	for _, tc := range []struct {
		seats []int
		want  int
	}{
		{[]int{1, 0, 0, 0}, 3},
		{[]int{1, 0, 0, 0, 0, 1, 0, 1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.seats), func(t *testing.T) {
			require.Equal(t, tc.want, maxDistToClosest(tc.seats))
		})
	}
}

func maxDistToClosest(seats []int) int {
	dist := make([]int, len(seats))
	curDist := math.MaxInt32 - len(seats)
	for i := range seats {
		if seats[i] == 1 {
			curDist = 0
		}
		dist[i] = curDist
		curDist++
	}

	n := len(seats)
	maxDist := 0
	curDist = dist[n-1]
	for i := n - 1; i >= 0; i-- {
		if seats[i] == 1 {
			curDist = 0
		}
		maxDist = max(maxDist, min(dist[i], curDist))
		curDist++
	}

	return maxDist
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
