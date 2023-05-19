package p1247minimumswapstomakestringsequal

func minimumSwap(s1 string, s2 string) int {
	// Count pairs
	var count [2][2]int
	for i := range s1 {
		count[s1[i]-'x'][s2[i]-'x']++
	}

	// There are two kinds of pairs that are interesting: xy and yx
	//
	// xy can match with xy to complete two pairs and vice versa for yx+yx
	//
	// If the number of xy or yx pairs is uneven, then one pair can be flipped
	// so that there is an even amount.
	//
	// Since pairs are always matched two-by-two, no solution can exist if the
	// total number of xy and yx pairs is uneven.
	if (count[0][1]+count[1][0])&1 == 1 {
		return -1
	}

	// The total number of swaps is equal to the number of pairings of xy+xy and
	// yx+yx, plus a single flip of an xy or yx to make the number of pairings
	// even.
	return (count[0][1]+count[1][0])/2 + (count[0][1]|count[1][0])&1
}
