package p2287rearrangecharacterstomaketargetstring

import "math"

func rearrangeCharacters(s string, target string) int {
	var freq [26]int
	for _, ch := range target {
		freq[ch-'a']++
	}
	var freq2 [26]int
	for _, ch := range s {
		freq2[ch-'a']++
	}
	maxMultiple := math.MaxInt32
	for ch, c := range freq {
		if c == 0 {
			continue
		}
		maxMultiple = min(maxMultiple, freq2[ch]/c)
	}
	return maxMultiple
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
