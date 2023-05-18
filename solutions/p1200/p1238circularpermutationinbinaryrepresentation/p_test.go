package p1238circularpermutationinbinaryrepresentation

func circularPermutation(n int, start int) []int {
	// Generate a list of indices to flip.
	// In order for permutation at index 1, 3, 7, etc to have exactly one bit
	// flip compared to start, then we must introduce a new flip index on those
	// indices. We must also un-flip that bit in an unoccupied location in the
	// sequence. This gives us:
	// Flip rightmost bit in perm 1, 3, 5, ...
	// Flip second-to-rightmost bit in perm 2, 6, 10
	// Then 4, 12, 20, etc.
	flips := make([]int, (1 << n))
	for x := 0; (1 << x) < (1 << n); x++ {
		d := 1 << (x + 1)
		for y := 1 << x; y < (1 << n); y += d {
			flips[y] = x
		}
	}
	res := make([]int, (1 << n))
	curr := start
	res[0] = start
	for i := 1; i < (1 << n); i++ {
		curr ^= 1 << flips[i]
		res[i] = curr
	}
	return res
}
