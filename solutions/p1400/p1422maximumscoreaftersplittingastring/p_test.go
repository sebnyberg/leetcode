package p1422maximumscoreaftersplittingastring

func maxScore(s string) int {
	var ones int
	for _, ch := range s {
		ones += int(ch - '0')
	}
	var res int
	var zeroes int
	for i := 0; i < len(s)-1; i++ {
		zeroes += 1 - int(s[i]-'0')
		ones -= int(s[i] - '0')
		res = max(res, ones+zeroes)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
