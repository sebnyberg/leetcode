package p1456maximumnumberofvowelsinasubstringofagivenlength

import "strings"

func maxVowels(s string, k int) int {
	var count int
	countVowel := func(r rune) int {
		if strings.ContainsRune("aeiou", r) {
			return 1
		}
		return 0
	}
	for i, ch := range s {
		count += countVowel(ch)
		if i >= k {
			count -= countVowel(rune(s[i-k]))
		}
	}
}
