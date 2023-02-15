package p1156swapforlongestrepeatedcharactersubstring

func maxRepOpt1(text string) int {
	// Slide a window over the text, counting any letters outside the window.
	//
	// The basic idea is that any valid sequence of repeated letters must have a
	// starting character. This character is the "left edge" of the sequence.
	//
	// At the moment that there are more than 1 characters that is not the same
	// as the starting character, the window is invalidated and a new starting
	// character must be found.
	var count [26]int
	for i := range text {
		count[text[i]-'a']++
	}

	var wcount [26]int
	var l int

	var res int
	for r := 0; r < len(text); r++ {
		wcount[text[r]-'a']++

		// While there is more than one character inequal to the starting
		// character, move the left pointer.
		for (r-l+1)-wcount[text[l]-'a'] > 1 {
			wcount[text[l]-'a']--
			l++
		}

		// Consider the 'a's in:
		//
		// 1. "abaaccc"
		// 2. "aaabccc"
		// 3. "abaacca"
		//
		// In case 1 and 2, the result is simply the count of 'a's in the window.
		// In case 3, we can borrow an external 'a' to get an extra letter.
		//
		nrepeat := wcount[text[l]-'a']
		if nrepeat < count[text[l]-'a'] {
			// can borrow a letter from somewhere else
			nrepeat++
		}
		res = max(res, nrepeat)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
