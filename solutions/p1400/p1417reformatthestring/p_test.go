package p1417reformatthestring

func reformat(s string) string {
	var letters []byte
	var digits []byte
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			letters = append(letters, byte(ch))
		} else {
			digits = append(digits, byte(ch))
		}
	}
	if abs(len(letters)-len(digits)) > 1 {
		return ""
	}
	var i int
	if len(digits) > len(letters) {
		i++
	}
	res := make([]byte, len(s))
	for k := 0; k < len(s); k++ {
		if i == 0 {
			res[k] = letters[k/2]
		} else {
			res[k] = digits[k/2]
		}
		i = (i + 1) % 2
	}
	return string(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
