package p2685countthenumberofcompletecomponents

func countCompleteComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)
	indeg := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		ra := find(parent[a])
		parent[a] = ra
		return ra
	}
	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra != rb {
			if size[ra] < size[rb] {
				ra, rb = rb, ra
			}
			parent[rb] = ra
			size[ra] += size[rb]
		}
	}
	for _, e := range edges {
		a := e[0]
		b := e[1]
		union(a, b)
		indeg[a]++
		indeg[b]++
	}
	for i := range parent {
		find(i)
	}
	ok := make(map[int]bool)
	for i := range parent {
		ok[parent[i]] = true
	}
	for i := range parent {
		if indeg[i]+1 != size[parent[i]] {
			ok[parent[i]] = false
		}
	}
	var res int
	for _, b := range ok {
		if b {
			res++
		}
	}
	return res
}
