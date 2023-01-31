package p1626bestteamwithnoconflicts

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n+1)
	}
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		if ages[idx[i]] == ages[idx[j]] {
			return scores[idx[i]] < scores[idx[j]]
		}
		return ages[idx[i]] < ages[idx[j]]
	})
	dp := make([]int, n)
	var res int
	for j := range idx {
		jj := idx[j]
		dp[jj] = scores[jj]
		for i := 0; i < j; i++ {
			ii := idx[i]
			if scores[ii] <= scores[jj] {
				dp[jj] = max(dp[jj], scores[jj]+dp[ii])
			}
		}
		res = max(res, dp[jj])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
