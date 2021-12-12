package p2101detonatethemaximumbombs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumDetonation(t *testing.T) {
	for _, tc := range []struct {
		bombs [][]int
		want  int
	}{
		{[][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}, 5},
		{[][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}, 5},
		{[][]int{{4, 4, 3}, {4, 4, 3}}, 2},
		{[][]int{{2, 1, 3}, {6, 1, 4}}, 2},
		{[][]int{{1, 1, 5}, {10, 10, 5}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bombs), func(t *testing.T) {
			require.Equal(t, tc.want, maximumDetonation(tc.bombs))
		})
	}
}

func maximumDetonation(bombs [][]int) int {
	// For each bomb, create an edge if it would result in detonating another bomb
	// upon exploding.
	n := len(bombs)

	// within returns true if b is within a
	within := func(a, b []int) bool {
		x1, y1, r := a[0], a[1], a[2]
		x2, y2 := b[0], b[1]
		dx := x2 - x1
		dy := y2 - y1
		return dx*dx+dy*dy <= r*r
	}

	// Create a directed graph
	adj := make([][]int, n)
	for i := range bombs {
		for j := i + 1; j < n; j++ {
			if within(bombs[i], bombs[j]) {
				adj[i] = append(adj[i], j)
			}
			if within(bombs[j], bombs[i]) {
				adj[j] = append(adj[j], i)
			}
		}
	}

	// Detonate bombs one by one
	var maxCount int
	cur := make([]int, 1, 100)
	for i := 0; i < n; i++ {
		var seen [100]bool
		cur = cur[:1]
		cur[0] = i
		seen[i] = true
		for i := 0; i < len(cur); i++ {
			for _, nei := range adj[cur[i]] {
				if seen[nei] {
					continue
				}
				seen[nei] = true
				cur = append(cur, nei)
			}
		}
		maxCount = max(maxCount, len(cur))
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
