package p2559countvowelstringsinranges

import "strings"

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	pre := make([]int, n+1)
	isvowel := func(r rune) bool {
		return strings.ContainsRune("aeiou", r)
	}
	for i := range words {
		m := len(words[i])
		pre[i+1] = pre[i]
		if isvowel(rune(words[i][0])) && isvowel(rune(words[i][m-1])) {
			pre[i+1]++
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = pre[q[1]+1] - pre[q[0]]
	}
	return res
}
