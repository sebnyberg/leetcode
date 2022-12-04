package p2492minimumscoreofapathbetweentwocities

import (
	"math"
	"sort"
)

func minScore(n int, roads [][]int) int {
	adj := make([][]int, n+1)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], r[1])
		adj[r[1]] = append(adj[r[1]], r[0])
	}
	sort.Slice(roads, func(i, j int) bool {
		return roads[i][2] < roads[j][2]
	})
	minPath := make([]int, n+1)
	for i := range minPath {
		minPath[i] = math.MaxInt32
	}
	curr := []int{}
	next := []int{}
	for _, e := range roads {
		if minPath[e[0]] != math.MaxInt32 {
			continue
		}
		var sawNodes [2]bool
		curr = curr[:0]
		curr = append(curr, e[0])
		if e[0] == 1 {
			sawNodes[0] = true
		}
		if e[0] == n {

			sawNodes[1] = true
		}
		minPath[e[0]] = e[2]
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, y := range adj[x] {
					if minPath[y] != math.MaxInt32 {
						continue
					}
					if y == 1 {
						sawNodes[0] = true
					}
					if y == n {
						sawNodes[1] = true
					}
					minPath[y] = e[2]
					next = append(next, y)
				}

			}
			curr, next = next, curr
		}
		if sawNodes[0] || sawNodes[1] {
			if !sawNodes[0] || !sawNodes[1] {

				return -1
			}
			return e[2]
		}
	}
	return -1
	// return minPath[find(1)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
