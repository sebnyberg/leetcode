package p2531makenumberofdistinctcharactersequal

func isItPossible(word1 string, word2 string) bool {
	var count1 [26]int
	for _, x := range word1 {
		x -= 'a'
		count1[x]++
	}
	var count2 [26]int
	for _, x := range word2 {
		x -= 'a'
		count2[x]++
	}
	var firstCount int
	for _, c := range count1 {
		if c > 0 {
			firstCount++
		}
	}
	var secondCount int
	for _, c := range count2 {
		if c > 0 {
			secondCount++
		}
	}
	for x := 0; x < 26; x++ {
		for y := 0; y < 26; y++ {
			if count1[x] == 0 || count2[y] == 0 {
				continue
			}
			// take x from w1 and put it in w2
			if x == y {
				if firstCount == secondCount {
					return true
				}
				continue
			}
			c1x := count1[x]
			c1y := count1[y]
			c2x := count2[x]
			c2y := count2[y]
			nextFirst := firstCount
			nextSecond := secondCount
			if c1x == 1 {
				nextFirst--
			}
			if c2x == 0 {
				nextSecond++
			}
			if c1y == 0 {
				nextFirst++
			}
			if c2y == 1 {
				nextSecond--
			}
			if nextFirst == nextSecond {
				return true
			}
		}
	}
	return false
}
