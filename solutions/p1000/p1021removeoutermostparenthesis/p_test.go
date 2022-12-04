package p1021removeoutermostparenthesis

func removeOuterParentheses(s string) string {
	// TL;DR: each parenthesis has a "depth", remove the smallest-depth
	// (shallowest) parenthesis.
	var res []rune
	var nparen int
	for _, ch := range s {
		if ch == '(' {
			nparen++
			if nparen > 1 {
				res = append(res, ch)
			}
		} else {
			nparen--
			if nparen > 0 {
				res = append(res, ch)
			}
		}
	}
	return string(res)
}
