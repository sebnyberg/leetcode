package p3133minimumarrayend

func minEnd(n int, x int) int64 {
	// For the AND operation to equal x, all bits in the sequence must contain the
	// bits in x.
	// Assuming that we start with x, which is the smallest possible number that
	// contains all the bits of x. Then we can add bits, and always OR with x to
	// ensure that the new number is large enough to contain x.
	// That should in the end yield the smallest possible final number.

	curr := x
	for i := 0; i < n-1; i++ {
		curr = (curr + 1) | x
	}
	return int64(curr)
}
