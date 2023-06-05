package p1761minimumdegreeofaconnectedtrioinagraph

import (
	"math"
	"sort"
)

func minTrioDegree(n int, edges [][]int) int {
	// For each edge, check if the pair of nodes connect to a shared, third
	// node. If they do, perform the check.
	//
	adj := make([][]int, n)
	for i, e := range edges {
		edges[i][0]--
		edges[i][1]--
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	// nei(i, j) returns true if the node i has the neighbour j
	// O(logn)
	nei := func(i, j int) bool {
		k := sort.SearchInts(adj[i], j)
		if k >= len(adj[i]) {
			return false
		}
		return adj[i][k] == j
	}

	for i := range adj {
		sort.Ints(adj[i])
	}

	res := math.MaxInt32
	for i := 0; i < n-2; i++ {
		startJ := sort.SearchInts(adj[i], i+1)
		if startJ >= len(adj[i]) {
			continue
		}
		for _, j := range adj[i][startJ:] {
			startK := sort.SearchInts(adj[j], j+1)
			if startK >= len(adj[j]) {
				continue
			}
			for _, k := range adj[j][startK:] {
				if !nei(i, k) {
					continue
				}
				deg := len(adj[i]) + len(adj[j]) + len(adj[k]) - 6
				res = min(res, deg)
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
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
