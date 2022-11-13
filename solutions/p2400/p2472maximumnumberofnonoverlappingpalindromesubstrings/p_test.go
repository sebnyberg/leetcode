package p2472maximumnumberofnonoverlappingpalindromesubstrings

func maxPalindromes(s string, k int) int {
	// For each position, for each palindrome prior to that position, add max
	// total count to dp
	var dp [2001]int
	ispalin := func(t string) bool {
		for l, r := 0, len(t)-1; l < r; l, r = l+1, r-1 {
			if t[l] != t[r] {
				return false
			}
		}
		return true
	}
	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1]
		for j := 0; j < i; j++ {
			if i-j >= k && ispalin(s[j:i]) {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}
	return dp[len(s)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
