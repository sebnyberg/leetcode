package p2761primepairswithtargetsum

func findPrimePairs(n int) [][]int {
	// Luckily, there aren't that many primes.
	sieve := make([]bool, n+1)
	primes := []int{}
	prime := make(map[int]bool)
	for x := 2; x <= n; x++ {
		if sieve[x] {
			continue
		}
		prime[x] = true
		primes = append(primes, x)
		for y := x; y <= n; y += x {
			sieve[y] = true
		}
	}
	var res [][]int
	for _, x := range primes {
		y := n - x
		if y < x {
			break
		}
		if prime[y] {
			res = append(res, []int{x, y})
		}
	}
	return res
}
