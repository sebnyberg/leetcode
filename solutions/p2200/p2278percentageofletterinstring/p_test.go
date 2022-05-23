package p2278percentageofletterinstring

func percentageLetter(s string, letter byte) int {
	var count int
	for _, ch := range s {
		if ch == rune(letter) {
			count++
		}
	}
	return (count * 100) / len(s)
}
