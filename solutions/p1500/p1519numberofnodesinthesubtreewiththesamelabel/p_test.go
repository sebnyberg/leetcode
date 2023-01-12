package p1519numberofnodesinthesubtreewiththesamelabel

func countSubTrees(n int, edges [][]int, labels string) []int {
	// Each node i sets res[i] and returns a [26]int count
	res := make([]int, n)
	adj := make([][]int, n)
	for _, e := range edges {
		a := e[0]
		b := e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	visit(adj, res, labels, -1, 0)
	return res
}

func visit(adj [][]int, res []int, labels string, parent, i int) [26]int {
	var count [26]int
	count[labels[i]-'a']++
	for _, v := range adj[i] {
		if v == parent {
			continue
		}
		for ch, c := range visit(adj, res, labels, i, v) {
			count[ch] += c
		}
	}
	res[i] = count[labels[i]-'a']
	return count
}
