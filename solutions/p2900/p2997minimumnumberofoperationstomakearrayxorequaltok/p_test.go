package p2997minimumnumberofoperationstomakearrayxorequaltok

import "math/bits"

func minOperations(nums []int, k int) int {
	// XOR zero means an even count of bits.
	// If we XOR all numbers, we get the odd count of bits. We can then xor that
	// with k to find the diff, and count bits.
	for _, x := range nums {
		k ^= x
	}
	return bits.OnesCount(uint(k))
}
