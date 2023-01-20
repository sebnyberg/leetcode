package p1

func defangIPaddr(address string) string {
	var res []byte
	for _, ch := range address {
		if ch == '.' {
			res = append(res, "[.]"...)
		} else {
			res = append(res, byte(ch))
		}
	}
	return string(res)
}
