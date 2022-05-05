package p1704determineifstringhalvesarealike

func halvesAreAlike(s string) bool {
	n := len(s)
	return countVowels(s[:n/2]) == countVowels(s[n/2:])
}

var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func countVowels(s string) int {
	count := 0
	for i := range s {
		if _, exists := vowels[s[i]]; exists {
			count++
		}
	}
	return count
}
