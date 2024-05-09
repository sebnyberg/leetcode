package p1657determineiftwostringsareclose

import "sort"

func closeStrings(word1 string, word2 string) bool {
	doCount := func(a string) ([]int, int) {
		var bm int
		res := make([]int, 26)
		for _, ch := range a {
			res[ch-'a']++
			bm |= (1 << (ch - 'a'))
		}
		return res, bm
	}
	a, ba := doCount(word1)
	b, bb := doCount(word2)
	if ba != bb {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
