package p1370increasingdecreasingstring

func sortString(s string) string {
	var count [26]int
	var m int
	for _, ch := range s {
		count[ch-'a']++
		m++
	}
	var res []byte
	for m > 0 {
		for k := 0; k < 26; k++ {
			if count[k] > 0 {
				res = append(res, byte(k+'a'))
				count[k]--
				m--
			}
		}
		for k := 25; k >= 0; k-- {
			if count[k] > 0 {
				res = append(res, byte(k+'a'))
				count[k]--
				m--
			}
		}
	}
	return string(res)
}
