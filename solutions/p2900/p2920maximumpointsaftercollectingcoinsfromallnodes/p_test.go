package p2920maximumpointsaftercollectingcoinsfromallnodes

func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	seen := make([]bool, n)
	adj := make([][]int, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	res := dfs(seen, adj, coins, 0, k)
	return res[0]
}

func dfs(seen []bool, adj [][]int, coins []int, i, k int) [16]int {
	seen[i] = true

	var childSums [16]int
	for _, child := range adj[i] {
		if seen[child] {
			continue
		}
		childVal := dfs(seen, adj, coins, child, k)
		for j, c := range childVal {
			childSums[j] += c
		}
	}
	var res [16]int

	for j := 0; j < 15; j++ {
		// a = choose not to reduce
		a := (coins[i] >> j) - k + childSums[j]

		// b = reduce current and child nodes
		b := (coins[i] >> (j + 1)) + childSums[j+1]

		res[j] = max(a, b)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
