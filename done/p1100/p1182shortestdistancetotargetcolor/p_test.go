package p1182shortestdistancetotargetcolor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestDistanceColor(t *testing.T) {
	for _, tc := range []struct {
		colors  []int
		queries [][]int
		want    []int
	}{
		{[]int{1, 1, 2, 1, 3, 2, 2, 3, 3}, [][]int{{1, 3}, {2, 2}, {6, 1}}, []int{3, 0, 3}},
		{[]int{1, 2}, [][]int{{0, 3}}, []int{-1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.colors), func(t *testing.T) {
			require.Equal(t, tc.want, shortestDistanceColor(tc.colors, tc.queries))
		})
	}
}

const maxLen = 50000

func shortestDistanceColor(colors []int, queries [][]int) []int {
	n := len(colors)
	dists := make([][3]uint16, n)

	lastSeen := []uint16{maxLen, maxLen, maxLen}

	// Forward pass
	for i, col := range colors {
		for j := range lastSeen {
			lastSeen[j]++
		}
		lastSeen[col-1] = 0
		for j := 0; j < 3; j++ {
			dists[i][j] = lastSeen[j]
		}
	}

	// Backward pass
	for i := n - 2; i >= 0; i-- {
		for j := range lastSeen {
			lastSeen[j]++
		}
		lastSeen[colors[i]-1] = 0
		for j := 0; j < 3; j++ {
			dists[i][j] = min(dists[i][j], lastSeen[j])
		}
	}

	results := make([]int, len(queries))
	for i, query := range queries {
		d := dists[query[0]][query[1]-1]
		if d >= maxLen {
			results[i] = -1
		} else {
			results[i] = int(d)
		}
	}

	return results
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
