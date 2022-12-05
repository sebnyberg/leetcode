package p1025divisorgame

func divisorGame(n int) bool {
	mem := make(map[int]bool)
	return dfs(mem, n)
}

func dfs(mem map[int]bool, n int) bool {
	if v, exists := mem[n]; exists {
		return v
	}
	if n == 1 {
		return false
	}
	for x := n - 1; x >= 1; x-- {
		if n%x == 0 {
			if !dfs(mem, n-x) {
				mem[n] = true
				return true
			}
		}
	}
	mem[n] = false
	return false
}
