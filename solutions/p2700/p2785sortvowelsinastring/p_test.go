package p2785sortvowelsinastring

import "sort"

func sortVowels(s string) string {
	vowels := []byte{}
	isvowel := [256]bool{}
	for _, ch := range []byte("aeiouAEIOU") {
		isvowel[ch] = true
	}
	for _, ch := range s {
		if isvowel[ch] {
			vowels = append(vowels, byte(ch))
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})
	res := []byte(s)
	var j int
	for i := range res {
		if isvowel[res[i]] {
			res[i] = vowels[j]
			j++
		}
	}
	return string(res)
}
