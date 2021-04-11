package p1816truncatesentence

func truncateSentence(s string, k int) string {
	k--
	for i, r := range s {
		if r == ' ' {
			if k == 0 {
				return s[:i]
			}
			k--
		}
	}
	return s
}
