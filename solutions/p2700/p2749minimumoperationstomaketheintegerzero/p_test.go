package p2749minimumoperationstomaketheintegerzero

import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
	// The key question is: could we remove a number in such a way that the
	// result becomes impossible, but would otherwise have been possible?

	// We are looking for a linear combination such that
	//
	// a * (2^0 + num2) + b * (2^1 + num2) + ... + z * (2^60 + num2) = 0
	//
	// This can be refactored as:
	//
	// a*num2 + b*num2 + ... + z*num2 + a*2^0 + b*2^1 + ... + z*2^60
	//
	// or
	// num2*(a+b+c+...+z) + a*2^0 + b*2^1 + ... + z*2^60
	//
	// For the factor related to num2, it does not matter where the coefficient
	// lies in terms of the combination of powers-of-two. Not sure if this matters.
	//
	// Aside for when num2 >= num1, when would there be no solution?
	//
	// num1 = 23, num2 = 8.
	//
	// This is not possible, because 23-8 = 15, which is not a power of 2. And
	// 23-16 = 7, which cannot be formed from two powers-of-two.
	//
	// When can a number be formed from a combination of powers-of-two? Well, if
	// the bitcount is greater than the count of powers, then it's impossible.
	// Also, if the number is smaller than the count of the powers, it's also
	// impossible.
	//
	if num2 >= num1 {
		return -1
	}

	x := num2
	for k := 1; ; k++ {
		y := num1 - x
		if y < k {
			if num1 > 0 {
				return -1
			}
			continue
		}
		if bits.OnesCount(uint(y)) <= k {
			return k
		}
		x += num2
	}
}
