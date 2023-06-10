package p1400constructkpalindromestrings

func canConstruct(s string, k int) bool {
	var freq [26]int
	for _, ch := range s {
		freq[ch-'a']++
	}
	var oddCount int
	for _, count := range freq {
		oddCount += count & 1
	}
	return oddCount <= k && k <= len(s)
}
