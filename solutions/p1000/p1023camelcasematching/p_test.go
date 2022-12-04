package p1023camelcasematching

import "unicode"

func camelMatch(queries []string, pattern string) []bool {
	validate := func(q string, p string) bool {
		var j int
		// Goal is to reach j == len(p)
		// Whenever we can match, we do
		for _, ch := range q {
			if j == len(p) {
				if unicode.IsUpper(ch) {
					return false
				}
				continue
			}
			if ch == rune(p[j]) {
				j++
				continue
			}
			if unicode.IsUpper(ch) {
				return false
			}
		}
		return j == len(p)
	}
	m := len(queries)
	res := make([]bool, m)
	for i, q := range queries {
		res[i] = validate(q, pattern)
	}
	return res
}
