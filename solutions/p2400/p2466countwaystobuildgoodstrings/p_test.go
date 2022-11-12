package p2466countwaystobuildgoodstrings

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	const mod = 1e9 + 7
	var res int
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero]
		}
		if i >= one {
			dp[i] += dp[i-one]
		}
		dp[i] %= mod
		if i >= low {
			res = (res + dp[i]) % mod
		}
	}
	return res
}
