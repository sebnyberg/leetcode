package p3110scoreofastring

func scoreOfString(s string) int {
	var res int
	for i := 1; i < len(s); i++ {
		res += abs(int(s[i]) - int(s[i-1]))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}