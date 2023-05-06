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
	var res int
	for i, ch := range s {
		count += countVowel(ch)
		if i >= k {
			count -= countVowel(rune(s[i-k]))
		}
		if i >= k-1 && count > res {
			res = count
		}
	}
	return res
}
