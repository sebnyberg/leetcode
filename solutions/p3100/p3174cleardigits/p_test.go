package p3174cleardigits

func clearDigits(s string) string {
	var res []byte
	isdigit := func(s byte) bool {
		return s >= '0' && s <= '9'
	}
	for i := range s {
		if isdigit(s[i]) && len(res) > 0 && !isdigit(res[len(res)-1]) {
			res = res[:len(res)-1]
			continue
		}
		res = append(res, s[i])
	}
	return string(res)
}
