package p2351firstlettertoappeartwice

func repeatedCharacter(s string) byte {
	// Use a bitmask to keep track of whether a character has
	// been seen before.
	var seen int
	for _, ch := range s {
		bit := (1 << (ch - 'a'))
		if seen&bit > 0 {
			return byte(ch)
		}
		seen |= bit
	}
	return 0
}
