package p1443minimumtimetocollectallapplesinatree

func minTime(n int, edges [][]int, hasApple []bool) int {
	adj := make(map[int][]int)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	seen := make([]bool, n)
	res := visit(seen, adj, hasApple, 0)
	if res > 0 {
		res -= 2
	}
	return res
}

func visit(seen []bool, adj map[int][]int, hasApple []bool, i int) int {
	var res int
	seen[i] = true
	for _, j := range adj[i] {
		if seen[j] {
			continue
		}
		res += visit(seen, adj, hasApple, j)
	}
	if hasApple[i] || res > 0 {
		res += 2
	}
	return res
}
