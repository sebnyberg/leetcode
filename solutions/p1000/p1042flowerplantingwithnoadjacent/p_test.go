package p1042flowerplantingwithnoadjacent

func gardenNoAdj(n int, paths [][]int) []int {
	// There is a solution, so we must only find a solution
	adj := make([][]int, n)
	for _, e := range paths {
		a := e[0] - 1
		b := e[1] - 1
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	res[0] = 0
	seen := make([]bool, n)
	for i := range adj {
		if res[i] == -1 {
			dfs(seen, res, adj, i)
		}
	}
	for i := range res {
		res[i]++
	}
	return res
}

func dfs(seen []bool, res []int, adj [][]int, i int) bool {
	// For each possible color for this node
	for k := 0; k < 4; k++ {
		// Check that it's ok
		for _, nei := range adj[i] {
			if res[nei] == k {
				goto cont
			}
		}
		res[i] = k
		// Color seems ok. Try it and move to all unvisited nodes
		for _, nei := range adj[i] {
			if res[nei] == -1 {
				if !dfs(seen, res, adj, nei) {
					goto cont
				}
			}
		}
		return true
	cont:
	}
	return false
}
