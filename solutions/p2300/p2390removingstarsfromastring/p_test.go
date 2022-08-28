package p2390removingstarsfromastring

func removeStars(s string) string {
	res := []rune{}
	for _, ch := range s {
		if ch == '*' && len(res) > 0 {
			res = res[:len(res)-1]
		}
		if ch != '*' {
			res = append(res, ch)
		}
	}
	return string(res)
}
