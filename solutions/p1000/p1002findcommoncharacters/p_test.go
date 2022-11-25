package p1002findcommoncharacters

func commonChars(words []string) []string {
	count := func(w string) [26]int {
		var res [26]int
		for _, ch := range w {
			res[ch-'a']++
		}
		return res
	}
	first := count(words[0])
	for _, w := range words {
		a := count(w)
		for i := range a {
			first[i] = min(first[i], a[i])
		}
	}
	var res []string
	for ch, n := range first {
		for n > 0 {
			res = append(res, string(byte(ch+'a')))
			n--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
