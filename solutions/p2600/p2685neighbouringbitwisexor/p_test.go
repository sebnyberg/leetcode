package p2685neighbouringbitwisexor

func doesValidArrayExist(derived []int) bool {
	// Rules:
	// XOR=1 => alternating
	// XOR=1 => same
	//
	// We can construct basically any solution to the first n-1 elements. If XOR
	// is 1, we alternate the original array, otherwise we keep elements the same.
	//
	// So we determine a kind of "same as first" or "unlike the first" requirement
	// and iterate over the array to see whether the condition holds for the final
	// element. That is, if the elements are 0....0 and the final element must be
	// unlike the first, then there can't be a valid original array, because the
	// only way to get 0 from XOR of the last and first element is for both to be
	// the same.

	n := len(derived)
	if n <= 1 {
		return derived[0] == 0
	}
	var differentFromFirst int
	for i := 0; i < n-1; i++ {
		// if derived[i] == 1, then we must flip differentFromFirst
		// 0 0 = 0
		// 0 1 = 1
		// 1 0 = 1
		// 1 1 = 0
		differentFromFirst ^= derived[i]
	}
	return derived[n-1] == differentFromFirst
}
