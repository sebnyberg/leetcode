package p0151reversewords

import "strings"

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	var trimmed strings.Builder
	var inspace bool
	for _, ch := range s {
		if ch == ' ' {
			if inspace {
				continue
			}
			inspace = true
			trimmed.WriteRune(ch)
		} else {
			inspace = false
			trimmed.WriteRune(ch)
		}
	}
	words := strings.Split(trimmed.String(), " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
