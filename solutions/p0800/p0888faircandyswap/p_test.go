package p0888faircandyswap

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	// Since an answer is guaranteed to exist, there must be a pair of numbers
	// that are not the same and that should be swapped.
	var delta int
	for _, x := range aliceSizes {
		delta += x
	}
	m := make(map[int]bool)
	for _, x := range bobSizes {
		delta -= x
		m[x] = true
	}
	delta /= 2
	// alice must trade delta to bob
	for _, x := range aliceSizes {
		if m[x-delta] {
			return []int{x, x - delta}
		}
	}
	return []int{0, 0}
}
