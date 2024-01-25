package p1624largestsubstringbetweentwoequalcharacters

import (
	"math"
)

func maxLengthBetweenEqualCharacters(s string) int {
	res := math.MinInt32
	var seen [26]int
	for i := range seen {
		seen[i] = math.MaxInt32
	}
	for i := range s {
		res = max(res, i-seen[s[i]-'a']-1)
		if seen[s[i]-'a'] == math.MaxInt32 {
			seen[s[i]-'a'] = i
		}
	}
	return max(-1, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
