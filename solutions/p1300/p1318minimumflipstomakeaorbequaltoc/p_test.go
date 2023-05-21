package p1318minimumflipstomakeaorbequaltoc

func minFlips(a int, b int, c int) int {
	var res int
	for a > 0 || b > 0 || c > 0 {
		if (a|b)&1 != c&1 {
			if c&1 == 0 {
				res += b&1 + a&1
			} else {
				res++
			}
		}
		a >>= 1
		b >>= 1
		c >>= 1
	}
	return res
}
