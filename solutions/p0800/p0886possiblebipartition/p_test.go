package p0886possiblebipartition

func possibleBipartition(n int, dislikes [][]int) bool {
	// This is a typical bipartite partitioning problem which can be verified
	// with coloring.
	color := make([]int, n+1)
	const (
		blue  = 0
		red   = 1
		unset = 2
	)
	for i := range color {
		color[i] = unset
	}
	adj := make([][]int, n+1)
	for _, d := range dislikes {
		adj[d[0]] = append(adj[d[0]], d[1])
		adj[d[1]] = append(adj[d[1]], d[0])
	}
	curr := []int{}
	next := []int{}
	for i := range color {
		if color[i] != unset {
			continue
		}
		color[i] = blue
		curr = curr[:0]
		curr = append(curr, i)
		for len(curr) > 0 {
			next = next[:0]
			for _, a := range curr {
				for _, b := range adj[a] {
					if color[b] == unset {
						color[b] = 1 - color[a]
						next = append(next, b)
						continue
					}
					if color[b] == color[a] {
						return false
					}
				}
			}
			curr, next = next, curr
		}
	}
	return true
}
