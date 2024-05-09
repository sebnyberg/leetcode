package p1143longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	// This is a classical DP problem
	// The easiest way to do this is using 2d-DP, but I'll just do 1D today
	// because I'm cocky
	n := len(text1)
	// m := len(text2)
	prev := make([]int, n+1)
	curr := make([]int, n+1)
	for i := range text2 {
		curr[0] = 0
		for j := 1; j <= n; j++ {
			curr[j] = max(curr[j-1], prev[j])
			if text1[j-1] == text2[i] {
				curr[j] = max(curr[j], prev[j-1]+1)
			}
		}
		curr, prev = prev, curr
	}
	return prev[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
