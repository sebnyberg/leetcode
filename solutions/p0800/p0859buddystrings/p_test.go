package p0859buddystrings

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		var charCount [26]int
		for i := range s {
			charCount[s[i]-'a']++
			if charCount[s[i]-'a'] > 1 {
				return true
			}
		}
		return false
	}

	// Not equal, try to find a valid pair swap
	bs := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == goal[i] {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[j] == goal[j] {
				continue
			}
			if s[j] != goal[i] || s[i] != goal[j] {
				continue
			}
			// Try to swap
			bs[i], bs[j] = bs[j], bs[i]
			if string(bs) == goal {
				return true
			}
			bs[i], bs[j] = bs[j], bs[i]
		}
	}
	return false
}
