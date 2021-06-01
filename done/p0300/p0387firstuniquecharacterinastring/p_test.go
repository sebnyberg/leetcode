package p0387firstuniquecharacterinastring

func firstUniqChar(s string) int {
	var count [26]int
	for _, ch := range s {
		count[ch-'a']++
	}
	for i, ch := range s {
		if count[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
