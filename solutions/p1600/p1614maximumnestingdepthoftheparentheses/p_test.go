package p1614maximumnestingdepthoftheparentheses

func maxDepth(s string) int {
	var paren int
	var maxParen int
	for _, ch := range s {
		if ch == '(' {
			paren++
		}
		if ch == ')' {
			paren--
		}
		maxParen = max(maxParen, paren)
	}
	return maxParen
}
