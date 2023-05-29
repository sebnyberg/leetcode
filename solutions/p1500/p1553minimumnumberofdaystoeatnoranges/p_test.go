package p1553minimumnumberofdaystoeatnoranges

func minDays(n int) int {
	mem := make(map[int]int)
	return dfs(mem, n)
}

func dfs(mem map[int]int, n int) int {
	if v, exists := mem[n]; exists {
		return v
	}
	if n <= 1 {
		return n
	}
	res := 1 + min(n%2+dfs(mem, n/2), n%3+dfs(mem, n/3))
	mem[n] = res
	return mem[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
