package p3223minimumlengthofstringafteroperations

func minimumLength(s string) int {
	// The order of processing does not matter,
	// so we can count each letter and calculate the result.
	var count [26]int
	for _, ch := range s {
		count[ch-'a']++
	}
	var res int
	for _, c := range count {
		if c&1 == 1 {
			res += 1
		} else {
			res += min(c, 2)
		}
	}
	return res
}
