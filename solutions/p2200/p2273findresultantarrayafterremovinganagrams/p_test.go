package p2273findresultantarrayafterremovinganagrams

func removeAnagrams(words []string) []string {
	var j int
	var prev [26]int
	res := make([]string, len(words))
	for _, w := range words {
		var freq [26]int
		for _, ch := range w {
			freq[ch-'a']++
		}
		if freq == prev {
			continue
		}
		res[j] = w
		j++
		prev = freq
	}
	res = res[:j]
	return res
}
