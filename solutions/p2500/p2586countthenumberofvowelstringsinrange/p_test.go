package p2586countthenumberofvowelstringsinrange

import "strings"

func vowelStrings(words []string, left int, right int) int {
	var count int
	for _, w := range words[left : right+1] {
		if strings.ContainsRune("aeiou", rune(w[0])) &&
			strings.ContainsRune("aeiou", rune(w[len(w)-1])) {
			count++
		}
	}
	return count
}
