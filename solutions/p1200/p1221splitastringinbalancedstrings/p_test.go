package p1221splitastringinbalancedstrings

func balancedStringSplit(s string) int {
	// Whenever there is a balanced substring, we should split immediately.
	// This is the way.
	//
	var l int
	var res int
	for i := range s {
		if s[i] == 'L' {
			l++
		} else {
			l--
		}
		if l == 0 {
			res++
		}
	}
	return res
}
