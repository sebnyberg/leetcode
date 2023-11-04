package p1503lastmomentbeforeallantsfalloutofaplank

func getLastMoment(n int, left []int, right []int) int {
	// "Changing directions" is actuall a no-op.
	//
	// So really what we're after is the longest distance from an ant's current
	// position and the end on the other side.
	//
	var res int
	for i := range left {
		res = max(res, left[i])
	}
	for i := range right {
		res = max(res, n-right[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
