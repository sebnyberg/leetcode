package p2930numberofstringswhichcanberearrangedtocontainsubstring

func stringCount(n int) int {
	// Any combination of letters that happen to contain leet will be fine.
	// What we really need to keep track of is the total count of l, e, and t.
	// Initially, the state is all zeros. For any given position, we can iterate
	// over the count of all possible prior states, then pick each letter and
	// update the state accordingly.
	var state [1 << 4]int
	const mod = 1e9 + 7

	const (
		lBit       = 1 << 0
		firstEBit  = 1 << 1
		secondEBit = 1 << 2
		tBit       = 1 << 3
	)
	state[0] = 1
	for i := 0; i < n; i++ {
		var nextState [1 << 4]int
		for x := 0; x < 1<<4; x++ {
			// Consider adding an 'l'
			if x&lBit == 0 {
				nextState[x|lBit] += state[x]
			} else {
				// Adding an l does not change anything
				nextState[x] += state[x]
			}

			if x&firstEBit == 0 {
				// Adding first e
				nextState[x|firstEBit] += state[x]
			} else if x&secondEBit == 0 {
				// Adding second e
				nextState[x|secondEBit] += state[x]
			} else {
				// Third, and so on
				nextState[x] += state[x]
			}

			// Consider adding the 't'
			if x&tBit == 0 {
				nextState[x|tBit] += state[x]
			} else {
				// Already added
				nextState[x] += state[x]
			}

			// Finally, all other letters simply keep state the same
			nextState[x] += state[x] * 23
			nextState[x] = nextState[x] % mod
		}
		state = nextState
	}
	return state[(1<<4)-1]
}
