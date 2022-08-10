package p0792numberofmatchingsubsequences

func numMatchingSubseq(s string, words []string) int {
	for _, ch := range s {
		for i := range words {
			if len(words[i]) > 0 && words[i][0] == byte(ch) {
				words[i] = words[i][1:]
			}
		}
	}
	var count int
	for _, w := range words {
		if len(w) == 0 {
			count++
		}
	}
	return count
}
