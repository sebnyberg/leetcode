package p1059allpathsfromsourceleadtodestination

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_leadsToDestination(t *testing.T) {
	for _, tc := range []struct {
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{1, [][]int{{0, 0}}, 0, 0, false},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1, 3, false},
		{3, [][]int{{0, 1}, {0, 2}}, 0, 2, false},
		{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 1}}, 0, 3, false},
		{4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 0, 3, true},
		{3, [][]int{{0, 1}, {1, 1}, {1, 2}}, 0, 2, false},
		{2, [][]int{{0, 1}, {1, 1}}, 0, 1, false},
	} {
		t.Run(fmt.Sprintf("%v/%+v", tc.n, tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, leadsToDestination(tc.n, tc.edges, tc.source, tc.destination))
		})
	}
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
	}
	if len(adj[destination]) > 0 {
		return false
	}

	colors := make([]byte, n)
	if hasCycle(adj, colors, n, source) {
		return false
	}

	// There are no cycles - now we can do DFS to validate that all paths lead
	// to the end.
	leadsToEnd := make([]byte, n)
	leadsToEnd[destination] = leadToEnd
	return allPathsLeadToEnd(adj, leadsToEnd, n, source)
}

const (
	unseen    = 0
	deadEnd   = 1
	leadToEnd = 2
)

func allPathsLeadToEnd(adj [][]int, leadsToEnd []byte, n, cur int) bool {
	if leadsToEnd[cur] != unseen {
		return leadsToEnd[cur] == leadToEnd
	}
	if len(adj[cur]) == 0 {
		leadsToEnd[cur] = deadEnd
		return false
	}
	for _, nei := range adj[cur] {
		if !allPathsLeadToEnd(adj, leadsToEnd, n, nei) {
			leadsToEnd[cur] = deadEnd
			return false
		}
	}
	leadsToEnd[cur] = leadToEnd
	return true
}

const (
	colorWhite = 0
	colorGrey  = 1
	colorBlack = 2
)

func hasCycle(adj [][]int, colors []byte, n, cur int) bool {
	colors[cur] = colorGrey
	for _, nei := range adj[cur] {
		if colors[nei] == colorGrey {
			return true
		}
		if hasCycle(adj, colors, n, nei) {
			return true
		}
	}
	colors[cur] = colorBlack
	return false
}
