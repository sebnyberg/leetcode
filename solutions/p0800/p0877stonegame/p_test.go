package p0877stonegame

func stoneGame(piles []int) bool {
	var dp [2][501][501]int
	n := len(piles)
	const playing, notPlaying = 0, 1
	for k := 1; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			j := i + k
			left := piles[i] + dp[notPlaying][i+1][j]
			right := piles[j-1] + dp[notPlaying][i][j-1]
			if left >= right {
				dp[playing][i][j] = left
				dp[notPlaying][i][j] = dp[playing][i+1][j]
			} else {
				dp[playing][i][j] = right
				dp[notPlaying][i][j] = dp[playing][i][j-1]
			}
		}
	}
	return dp[playing][0][n] > dp[notPlaying][0][n]
}
