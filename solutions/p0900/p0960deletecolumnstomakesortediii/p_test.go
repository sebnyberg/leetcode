package p0960deletecolumnstomakesortediii

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 1; j < n; j++ {
		dp[j] = 1
		for k := 0; k < j; k++ {
			for i := range strs {
				if strs[i][j] < strs[i][k] {
					goto cont
				}
			}
			dp[j] = max(dp[j], dp[k]+1)
		cont:
		}
	}
	res := n
	for _, v := range dp {
		res = min(res, n-v)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
