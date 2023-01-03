package p1071greatestcommondivisorofstrings

func gcdOfStrings(str1 string, str2 string) string {
	if len(str2) < len(str1) {
		str1, str2 = str2, str1
	}
	if str2[:len(str1)] != str1 {
		return ""
	}
	for n := len(str1); n > 0; n-- {
		if len(str2)%n != 0 || len(str1)%n != 0 {
			continue
		}
		for i := n; i < len(str2); i += n {
			if str2[i:i+n] != str2[:n] {
				goto notOk
			}
		}
		return str2[:n]
	notOk:
	}
	return ""
}
