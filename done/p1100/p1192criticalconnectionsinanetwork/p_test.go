package p1192criticalconnectionsinanetwork

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_criticalConnections(t *testing.T) {
	for _, tc := range []struct {
		n           int
		connections [][]int
		want        [][]int
	}{
		{4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}, [][]int{{1, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, criticalConnections(tc.n, tc.connections))
		})
	}
}

func criticalConnections(n int, connections [][]int) [][]int {
	adj := make([][]int, n)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	t := &Tarjan{
		time:       1,
		timestamps: make([]int, n),
		low:        make([]int, n),
	}
	t.dfs(adj, 0, -1)
	return t.bridges
}

type Tarjan struct {
	time       int
	timestamps []int
	low        []int
	bridges    [][]int
}

func (t *Tarjan) dfs(adj [][]int, cur, parent int) {
	t.timestamps[cur] = t.time
	t.low[cur] = t.time
	t.time++

	for _, nei := range adj[cur] {
		if nei == parent { // avoid checking the parent
			continue
		}
		// If the nei[ghbour] node has not been seen before
		// Continue SCC search and set the current low-link
		// value to that node's low-link value.
		if t.timestamps[nei] == 0 {
			t.dfs(adj, nei, cur)
			t.low[cur] = min(t.low[cur], t.low[nei])
			if t.timestamps[cur] < t.low[nei] {
				t.bridges = append(t.bridges, []int{cur, nei})
			}
		} else { // Seen before and part of this scc, use timestamp as lowlink
			t.low[cur] = min(t.low[cur], t.timestamps[nei])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
