package p0917reverseonlyletters

func reverseOnlyLetters(s string) string {
	isalpha := func(b byte) bool {
		return (b >= 'a' && b <= 'z') ||
			(b >= 'A' && b <= 'Z')
	}
	bs := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		// find left
		for l < r && !isalpha(s[l]) {
			l++
		}
		// find right
		for l < r && !isalpha(s[r]) {
			r--
		}
		// swap
		if l != r {
			bs[l], bs[r] = bs[r], bs[l]
		}
		l++
		r--
	}
	return string(bs)
}
