package p2928distributecandiesamongchildreni

func distributeCandies(n int, limit int) int {
	var curr [51]int
	curr[n] = 1
	for i := 0; i < 3; i++ {
		var next [51]int
		for k := 0; k <= limit; k++ {
			for j := 0; j+k <= n; j++ {
				next[j] += curr[j+k]
			}
		}
		curr = next
	}
	return curr[0]
}
