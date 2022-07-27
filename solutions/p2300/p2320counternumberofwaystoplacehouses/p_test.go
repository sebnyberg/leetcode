package p2320counternumberofwaystoplacehouses

func countHousePlacements(n int) int {
	// We can count configurations on one side only
	// then take the result * result for the actual result

	const mod = 1e9 + 7

	// dp[0] = number of configurations which have zero houses in the last slot
	// dp[1] = number of configurations which have a house in the last slot
	dp := [2]int{1, 1}
	for i := 2; i <= n; i++ {
		dp[0], dp[1] = (dp[1]+dp[0])%mod, dp[0]
	}

	m := (dp[0] + dp[1]) % mod
	return (m * m) % mod
}
