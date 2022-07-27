package p2316countunreachablepairsofnodesinanundirectedgraph

func countPairs(n int, edges [][]int) int64 {
	// Create adj list from edges
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	seen := make([]bool, n)
	curr := []int{}
	next := []int{}
	var res int
	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		// Start bfs
		curr := curr[:0]
		curr = append(curr, i)
		seen[i] = true
		count := 1
		for len(curr) > 0 {
			next = next[:0]
			for _, node := range curr {
				for _, nei := range adj[node] {
					if seen[nei] {
						continue
					}
					seen[nei] = true
					count++
					next = append(next, nei)
				}
			}
			curr, next = next, curr
		}
		res += (n - count) * count
	}

	return int64(res) / 2
}
