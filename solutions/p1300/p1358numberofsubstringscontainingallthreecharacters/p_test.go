package p1358numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) int {
	var l int
	var count [3]int
	var m int
	var res int
	for i := range s {
		ch := int(s[i] - 'a')
		count[ch]++
		if count[ch] == 1 {
			m++
		}
		if m < 3 {
			continue
		}
		for count[s[l]-'a'] > 1 {
			count[s[l]-'a']--
			l++
		}
		res += l + 1
	}
	return res
}
