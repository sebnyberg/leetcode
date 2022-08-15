package p0819mostcommonword

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	r := strings.NewReplacer(
		"!", " ",
		"?", " ",
		"'", " ",
		"\"", " ",
		";", " ",
		",", " ",
		".", " ",
	)
	paragraph = strings.ToLower(paragraph)
	paragraph = r.Replace(paragraph)
	b := make(map[string]struct{})
	for _, bb := range banned {
		b[strings.ToLower(bb)] = struct{}{}
	}
	isBanned := func(s string) bool {
		_, exists := b[s]
		return exists
	}
	count := make(map[string]int)
	for _, w := range strings.Fields(paragraph) {
		if isBanned(w) {
			continue
		}
		count[w]++
	}
	var maxCount int
	var maxWord string
	for w, c := range count {
		if c > maxCount || c == maxCount && w < maxWord {
			maxCount = c
			maxWord = w
		}
	}
	return maxWord
}
