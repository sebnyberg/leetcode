package p1832checkifsentenceispangram

func checkIfPangram(sentence string) bool {
	var seen [26]bool
	for _, ch := range sentence {
		seen[ch-'a'] = true
	}
	for _, s := range seen {
		if !s {
			return false
		}
	}
	return true
}
