package p3042countprefixandsuffixpairsi

func countPrefixSuffixPairs(words []string) int {
	isPrefixAndSuffix := func(a, b string) bool {
		if len(a) > len(b) {
			return false
		}
		n := len(b)
		for i := range a {
			if b[i] != a[i] || b[n-len(a)+i] != a[i] {
				return false
			}
		}
		return true
	}
	var res int
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				res++
			}
		}
	}
	return res
}
