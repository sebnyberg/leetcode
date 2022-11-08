package p1544makethestringgreat

import "unicode"

func makeGood(s string) string {
	d := int('A' - 'a')
	for i := 0; i < len(s)-1; {
		if abs(int(s[i])-int(s[i+1])) == d {
			s = s[:i] + s[i+2:]
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return s
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
