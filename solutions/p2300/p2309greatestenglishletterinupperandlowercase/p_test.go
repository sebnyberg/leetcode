package p2309greatestenglishletterinupperandlowercase

import "unicode"

func greatestLetter(s string) string {
	var seen int
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			seen |= (1 << (26 + (ch - 'A')))
		} else {
			seen |= (1 << (ch - 'a'))
		}
	}
	for ch := 25; ch >= 0; ch-- {
		m := (1 << ch) | (1 << (26 + ch))
		if seen&m == m {
			return string(ch + 'A')
		}
	}
	return ""
}
