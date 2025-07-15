package p3136validword

import (
	"strings"
)

func isValid(word string) bool {
	var hasVowel, hasConsonant bool
	for _, c := range word {
		switch {
		case c >= '0' && c <= '9':
			continue
		case c >= 'A' && c <= 'Z':
			c -= 'A' - 'a'
			fallthrough
		case c >= 'a' && c <= 'z':
			isVowel := strings.ContainsRune("aeiou", c)
			hasVowel = hasVowel || isVowel
			hasConsonant = hasConsonant || !isVowel
		default:
			return false
		}
	}
	return len(word) >= 3 && hasVowel && hasConsonant
}
