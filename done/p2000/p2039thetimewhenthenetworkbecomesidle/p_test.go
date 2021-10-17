package p2039thetimewhenthenetworkbecomesidle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_networkBecomesIdle(t *testing.T) {
	for _, tc := range []struct {
		edges    [][]int
		patience []int
		want     int
	}{
		{[][]int{{0, 1}, {1, 2}}, []int{0, 2, 1}, 8},
		{[][]int{{0, 1}, {0, 2}, {1, 2}}, []int{0, 10, 10}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, networkBecomesIdle(tc.edges, tc.patience))
		})
	}
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)

	// Create adj list
	adj := make([][]int, n)
	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// Perform BFS to find the distance to each edge
	dist := make([]int, n)
	seen := make([]bool, n)
	seen[0] = true
	dist[0] = 0
	nseen := 1
	steps := 1
	tovisit := []int{0}
	next := []int{}
	for nseen < n {
		next = next[:0]
		for _, cur := range tovisit {
			for _, nei := range adj[cur] {
				if !seen[nei] {
					seen[nei] = true
					dist[nei] = steps
					next = append(next, nei)
					nseen++
				}
			}
		}
		next, tovisit = tovisit, next
		steps++
	}

	var maxTime int
	for i := 1; i < len(dist); i++ {
		packetRoundtrip := dist[i] * 2
		lastPacketSent := ((dist[i]*2 - 1) / patience[i]) * patience[i]
		totalTime := lastPacketSent + packetRoundtrip
		maxTime = max(maxTime, totalTime)
	}

	return maxTime + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
