package p2368reachablenodeswithrestrictions

func reachableNodes(n int, edges [][]int, restricted []int) int {
	curr := []int{0}
	next := []int{}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	seen := make([]bool, n)
	for _, v := range restricted {
		seen[v] = true
	}
	seen[0] = true
	res := 1
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, nei := range adj[x] {
				if seen[nei] {
					continue
				}
				seen[nei] = true
				res++
				next = append(next, nei)
			}
		}
		curr, next = next, curr
	}
	return res
}
