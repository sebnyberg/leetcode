package p0983minimumcostfortickets

func mincostTickets(days []int, costs []int) int {
	var dp [366]int
	var j int
	for d := 1; d < 366; d++ {
		if d != days[j] {
			dp[d] = dp[d-1]
		} else {
			dp[d] = dp[d-1] + costs[0]
			dp[d] = min(dp[d], dp[max(0, d-7)]+costs[1])
			dp[d] = min(dp[d], dp[max(0, d-30)]+costs[2])
			j++
			if j == len(days) {
				break
			}
		}
	}
	return dp[days[len(days)-1]]
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
