package p0

import "sort"

func countPrefixes(words []string, s string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var count int
	for _, w := range words {
		if len(w) > len(s) {
			break
		}
		if s[:len(w)] == w {
			count++
		}
	}
	return count
}
