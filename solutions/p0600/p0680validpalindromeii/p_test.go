package p0680validpalindromeii

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			// Try to skip l or r
			return check(s[l+1:r+1]) || check(s[l:r])
		}
	}
	return true
}

func check(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
