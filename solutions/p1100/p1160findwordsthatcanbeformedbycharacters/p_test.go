package p1160findwordsthatcanbeformedbycharacters

func countCharacters(words []string, chars string) int {
	var count [26]int
	for _, c := range chars {
		count[c-'a']++
	}
	var res int
loop:
	for _, x := range words {
		var m [26]int
		for _, ch := range x {
			m[ch-'a']++
		}
		for i := range m {
			if count[i] < m[i] {
				continue loop
			}
		}
		res += len(x)
	}
	return res
}
