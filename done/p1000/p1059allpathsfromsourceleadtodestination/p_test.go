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
	indeg := make([]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		indeg[edge[1]]++
	}

	visited := make([]bool, n)
	tovisit := []int{}
	for idx, deg := range indeg {
		if deg == 0 {
			tovisit = append(tovisit, idx)
			visited[idx] = true
		}
	}

	// detect cycle using toposort
	next := []int{}
	for len(tovisit) > 0 {
		next = next[:0]
		for _, n := range tovisit {
			for _, nei := range adj[n] {
				indeg[nei]--
				if visited[nei] {
					continue
				}
				if indeg[nei] == 0 {
					visited[nei] = true
					next = append(next, nei)
				}
			}
		}
		tovisit, next = next, tovisit
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}

	// there are no cycles, perform DFS to check if all paths end in destination
	status := make([]byte, n)
	status[destination] = leadsToEnd
	dfs(adj, status, n, source, destination)
	for _, st := range status {
		if st != leadsToEnd {
			return false
		}
	}
	return true
}

const (
	notVisited byte = 0
	deadEnd    byte = 1
	leadsToEnd byte = 2
)

func dfs(adj [][]int, status []byte, n, cur, destination int) bool {
	if status[cur] != notVisited {
		return status[cur] == leadsToEnd
	}
	if len(adj[cur]) == 0 {
		status[cur] = deadEnd
		return false
	}
	ok := true
	for _, nei := range adj[cur] {
		if !dfs(adj, status, n, nei, destination) {
			ok = false
		}
	}
	if ok {
		status[cur] = leadsToEnd
	} else {
		status[cur] = deadEnd
	}
	return ok
}
