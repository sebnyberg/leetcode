package p2451oddstringdifference

func oddString(words []string) string {
	m := make(map[[20]int]int)
	calc := func(w string) [20]int {
		var res [20]int
		for i := 1; i < len(w); i++ {
			res[i-1] = int(w[i] - w[i-1])
		}
		return res
	}
	for _, w := range words {
		m[calc(w)]++
	}
	for _, w := range words {
		if m[calc(w)] == 1 {
			return w
		}
	}
	return ""
}
