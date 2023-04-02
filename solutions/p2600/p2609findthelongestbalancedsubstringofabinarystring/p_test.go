package p2609findthelongestbalancedsubstringofabinarystring

func findTheLongestBalancedSubstring(s string) int {
	var zeroes int
	var ones int
	var res int
	for _, ch := range s {
		if ch == '0' {
			if ones > 0 {
				ones = 0
				zeroes = 0
			}
			zeroes++
		} else {
			ones++
			res = max(res, min(ones, zeroes)*2)
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
