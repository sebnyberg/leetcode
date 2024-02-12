package p2939maximumxorproduct

func maximumXorProduct(a int64, b int64, n int) int {
	// For each bit, there are two options:
	// The bit in a and b are the same => pick the opposite bit for x
	// The bit in a and b are different => pick the opposite bit of the smallest
	// currently xor valued out of a and b.

	// Copy a/b and mask
	aa := a
	aa &^= (1 << n) - 1
	bb := b
	bb &^= (1 << n) - 1

	for x := n - 1; x >= 0; x-- {
		bit := int64(1 << x)
		if a&bit == b&bit {
			// easy pickings
			aa += bit
			bb += bit
		} else {
			if aa < bb {
				aa += bit
			} else {
				bb += bit
			}
		}
	}
	const mod = 1e9 + 7
	return int((aa%mod)*(bb%mod)) % mod
}
