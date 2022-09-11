package p2405optimalpartitionofstring

func partitionString(s string) int {
	var seen int
	count := 1
	for _, ch := range s {
		bit := int(1 << (ch - 'a'))
		if seen&bit > 0 {
			seen = 0
			count++
		}
		seen |= bit
	}
	return count
}
