package p0935knightdialer

func knightDialer(n int) int {
	mvts := [][2]int{
		{-2, -1}, {-1, -2},
		{2, -1}, {1, -2},
		{-2, 1}, {-1, 2},
		{2, 1}, {1, 2},
	}
	ok := func(i, j int) bool {
		if i == 3 {
			return j == 1
		}
		return i >= 0 && i <= 2 && j >= 0 && j <= 2
	}
	var dp [4][3]int
	for i := range dp {
		for j := range dp[i] {
			if ok(i, j) {
				dp[i][j] = 1
			}
		}
	}
	const mod = 1e9 + 7
	for k := 2; k <= n; k++ {
		var next [4][3]int
		for i := range dp {
			for j := range dp[i] {
				if !ok(i, j) {
					continue
				}
				for _, m := range mvts {
					ii := i + m[0]
					jj := j + m[1]
					if ok(ii, jj) {
						next[i][j] = (next[i][j] + dp[ii][jj]) % mod
					}
				}
			}
		}
		dp = next
	}
	var res int
	for i := range dp {
		for j := range dp[i] {
			res = (res + dp[i][j]) % mod
		}
	}
	return res
}
