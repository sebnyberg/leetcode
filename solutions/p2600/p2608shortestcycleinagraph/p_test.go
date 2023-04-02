package p2608shortestcycleinagraph

import "math"

func findShortestCycle(n int, edges [][]int) int {
	// Perform DFS and mark each node with an index.
	// The cycle length is the difference in current and marked index
	adj := make([][]int, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	res := math.MaxInt32
	for i := range adj {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		parent := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}
		dist[i] = 0
		q := []int{i}
		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			for _, y := range adj[x] {
				if dist[y] == math.MaxInt32 {
					dist[y] = 1 + dist[x]
					parent[y] = x
					q = append(q, y)
					continue
				}
				if parent[y] != x && parent[x] != y {
					res = min(res, dist[x]+dist[y]+1)
				}
			}
		}
	}
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
