package p1466reorderroutestomakeallpathsleadtothecityzero

func minReorder(n int, connections [][]int) int {
	adj := make([][]int, n)
	out := make([][]bool, n)
	for _, c := range connections {
		a := c[0]
		b := c[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
		out[a] = append(out[a], true)
		out[b] = append(out[b], false)
	}
	curr := []int{0}
	next := []int{}
	seen := make([]bool, n)
	seen[0] = true
	var res int
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for j, y := range adj[x] {
				if seen[y] {
					continue
				}
				seen[y] = true
				if out[x][j] {
					res++
				}
				next = append(next, y)
			}
		}

		curr, next = next, curr
	}
	return res
}
