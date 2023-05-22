package p1406stonegameiii

import "math"

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = math.MinInt32
	}
	// m[i] = maximum number of points for the current player given that the
	// front-stone is the i'th in stoneValue.
	res := dfs(mem, stoneValue, 0, len(stoneValue))
	if res < 0 {
		return "Bob"
	} else if res == 0 {
		return "Tie"
	}
	return "Alice"
}

func dfs(mem []int, stoneValue []int, i, n int) int {
	if i == n {
		return 0
	}
	if mem[i] != math.MinInt32 {
		return mem[i]
	}
	res := math.MinInt32
	var sum int
	for k := 1; k <= min(n-i, 3); k++ {
		sum += stoneValue[i+k-1]
		res = max(res, sum-dfs(mem, stoneValue, i+k, n))
	}
	mem[i] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
