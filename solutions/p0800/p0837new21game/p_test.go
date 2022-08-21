package p3

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k+maxPts {
		return 1
	}
	dp := make([]float64, n+1)
	dp[0] = 1
	windowSum := 1.0
	var res float64
	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i < k {
			windowSum += dp[i]
		} else {
			res += dp[i]
		}
		if i >= maxPts {
			windowSum -= dp[i-maxPts]
		}
	}
	return 0
}
