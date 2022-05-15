package p1759countnumberofhomgenoussubstrings

func countHomogenous(s string) int {
	var count int
	var res int
	for i := range s {
		if i > 0 && s[i-1] == s[i] {
			count++
		} else {
			count = 1
		}
		res = (res + count) % (1e9 + 7)
	}
	return res
}
