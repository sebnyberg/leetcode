package p2496maximumvalueofastringinanarray

func maximumValue(strs []string) int {
	isnum := func(s string) bool {
		for i := range s {
			if s[i] < '0' || s[i] > '9' {
				return false
			}
		}
		return true
	}

	var res int
	for _, s := range strs {
		if isnum(s) {
			var x int
			for i := range s {
				x *= 10
				x += int(s[i] - '0')
			}
			res = max(res, x)
		} else {
			res = max(res, len(s))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
