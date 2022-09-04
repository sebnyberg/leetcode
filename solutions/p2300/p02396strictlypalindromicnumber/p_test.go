package p02396strictlypalindromicnumber

func isStrictlyPalindromic(n int) bool {
	// Let's consider brute force
	// If we start with high bases, we are more likely to invalidate the number
	// quickly.
	isPalindrome := func(bs []int) bool {
		for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
			if bs[l] != bs[r] {
				return false
			}
		}
		return true
	}
	buf := []int{}
	check := func(b int) bool {
		buf = buf[:0]
		x := n
		for x > 0 {
			buf = append(buf, x%b)
			x /= b
		}
		return isPalindrome(buf)
	}
	for b := n - 2; b >= 2; b-- {
		if !check(b) {
			return false
		}
	}
	return true
}
