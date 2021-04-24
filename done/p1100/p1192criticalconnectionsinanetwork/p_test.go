package p1192criticalconnectionsinanetwork

import (
	"fmt"
	"math"
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
		bridges:    make([][]int, 0),
	}
	t.dfs(adj, n, 0, -1)
	return t.bridges
}

type Tarjan struct {
	time       int
	timestamps []int
	bridges    [][]int
}

func (t *Tarjan) dfs(adj [][]int, n, cur, parent int) int {
	if t.timestamps[cur] != 0 {
		return t.timestamps[cur]
	}
	t.timestamps[cur] = t.time
	t.time++

	minTimestamp := math.MaxInt32
	for _, nei := range adj[cur] {
		if nei == parent { // avoid checking the parent
			continue
		}
		neiTs := t.dfs(adj, n, nei, cur)
		minTimestamp = min(minTimestamp, neiTs)
	}
	// if minTimestamp >= t.timestamps[cur] && parent >= 0 {
	// 	t.bridges = append(t.bridges, []int{parent, cur})
	// }
	return min(t.timestamps[cur], minTimestamp)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
