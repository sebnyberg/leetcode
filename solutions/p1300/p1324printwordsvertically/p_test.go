package p1324printwordsvertically

import "strings"

func printVertically(s string) []string {
	words := strings.Split(s, " ")
	var maxLen int
	for _, w := range words {
		maxLen = max(maxLen, len(w))
	}
	res := make([][]byte, maxLen)
	for _, w := range words {
		for i := 0; i < maxLen; i++ {
			if i >= len(w) {
				res[i] = append(res[i], ' ')
			} else {
				res[i] = append(res[i], w[i])
			}
		}
	}
	ret := make([]string, len(res))
	for i := range res {
		ret[i] = strings.TrimRight(string(res[i]), " ")
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
