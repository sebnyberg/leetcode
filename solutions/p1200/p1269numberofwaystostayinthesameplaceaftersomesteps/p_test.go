package p1269numberofwaystostayinthesameplaceaftersomesteps

const mod = 1e9 + 7

func numWays(steps int, arrLen int) int {
	// Just simulating the steps is easy enough. There are only up to 500 steps,
	// which means that we may walk at most 250 steps to the right before
	// returning again. All other locations are irrelevant.
	// Because off-by-one is annoying, we may use the entire step length. It
	// won't impact the runtime performance much anyway.
	n := min(arrLen, steps)
	curr := make([]int, n)
	next := make([]int, n)
	curr[0] = 1
	for x := 0; x < steps; x++ {
		for i := range next {
			x := curr[i]
			if i > 0 {
				x += curr[i-1]
			}
			if i < n-1 {
				x += curr[i+1]
			}
			next[i] = x % mod
		}
		curr, next = next, curr
	}
	return curr[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
