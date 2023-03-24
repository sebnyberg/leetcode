package p1466reorderroutestomakeallpathsleadtothecityzero

func minReorder(n int, connections [][]int) int {
	adj := make([][]int, n)
	for _, c := range connections {
		a := c[0]
		b := c[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], -a)
	}
	curr := []int{0}
	next := []int{}
	seen := make([]bool, n)
	seen[0] = true
	var res int
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, y := range adj[x] {
				out := y < 0
				if out {
					y = -y
				}
				if seen[y] {
					continue
				}
				seen[y] = true
				if !out {
					res++
				}
				next = append(next, y)
			}
		}

		curr, next = next, curr
	}
	return res
}
