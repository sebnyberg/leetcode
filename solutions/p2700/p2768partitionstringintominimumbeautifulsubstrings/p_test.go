package p2768partitionstringintominimumbeautifulsubstrings

import (
	"fmt"
	"math"
)

func minimumBeautifulSubstrings(s string) int {
	want := make(map[string]bool)
	want["1"] = true
	for x := 5; ; x *= 5 {
		t := fmt.Sprintf("%b", x)
		if len(t) > len(s) {
			break
		}
		want[t] = true
	}
	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := range s {
		for w := range want {
			if len(w) > i+1 {
				continue
			}
			if s[i-len(w)+1:i+1] != w {
				continue
			}
			dp[i+1] = min(dp[i+1], dp[i-len(w)+1]+1)
		}
	}
	if dp[n] >= math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
