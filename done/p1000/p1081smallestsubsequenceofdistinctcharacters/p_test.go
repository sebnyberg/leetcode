package p1081smallestsubsequenceofdistinctcharacters

func smallestSubsequence(s string) string {
	var letterCount [26]int
	for _, ch := range s {
		letterCount[ch-'a']++
	}
	letters := make([]rune, 0)
	n := 0
	var hasLetter [26]bool
	for _, ch := range s {
		letterCount[ch-'a']--
		if hasLetter[ch-'a'] {
			continue
		}
		for len(letters) > 0 && letters[n-1] > ch && letterCount[letters[n-1]-'a'] > 0 {
			hasLetter[letters[n-1]-'a'] = false
			letters = letters[:n-1] // pop
			n--
		}
		hasLetter[ch-'a'] = true
		n++
		letters = append(letters, ch)
	}

	return string(letters)
}
