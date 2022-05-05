package p0582killprocess

func killProcess(pid []int, ppid []int, kill int) []int {
	adj := make(map[int][]int)
	for i, pp := range ppid {
		adj[pp] = append(adj[pp], pid[i])
	}

	seen := make(map[int]struct{})
	seen[kill] = struct{}{}
	curr := []int{kill}
	res := []int{kill}
	next := []int{}
	for len(curr) > 0 {
		next = next[:0]
		for _, n := range curr {
			for _, nei := range adj[n] {
				if _, exists := seen[nei]; exists {
					continue
				}
				seen[nei] = struct{}{}
				next = append(next, nei)
				res = append(res, nei)
			}
		}

		curr, next = next, curr
	}

	return res
}
