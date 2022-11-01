package typex

import (
	"fmt"
	"math/bits"
)

// uint128 represents a uint128 using two uint64s.
//
// Most of this code was taken from net/netip/uint128
//
// When the methods below mention a bit number, bit 0 is the most
// significant bit (in hi) and bit 127 is the lowest (lo&1).
type uint128 struct {
	hi uint64
	lo uint64
}

// mask6 returns a uint128 bitmask with the bottommost n bits of a
// 128-bit number.
func mask6(n int) uint128 {
	return uint128{
		^uint64(0) >> (128 - n),
		^(^uint64(0) << n),
	}
}

// isZero reports whether u == 0.
//
// It's faster than u == (uint128{}) because the compiler (as of Go
// 1.15/1.16b1) doesn't do this trick and instead inserts a branch in
// its eq alg's generated code.
func (u uint128) isZero() bool { return u.hi|u.lo == 0 }

// and returns the bitwise AND of u and m (u&m).
func (u uint128) and(m uint128) uint128 {
	return uint128{u.hi & m.hi, u.lo & m.lo}
}

// xor returns the bitwise XOR of u and m (u^m).
func (u uint128) xor(m uint128) uint128 {
	return uint128{u.hi ^ m.hi, u.lo ^ m.lo}
}

// or returns the bitwise OR of u and m (u|m).
func (u uint128) or(m uint128) uint128 {
	return uint128{u.hi | m.hi, u.lo | m.lo}
}

// not returns the bitwise NOT of u.
func (u uint128) not() uint128 {
	return uint128{^u.hi, ^u.lo}
}

func (u uint128) sub(m uint64) uint128 {
	lo, borrow := bits.Sub64(u.lo, m, 0)
	return uint128{u.hi - borrow, lo}
}

func (u uint128) add(m uint64) uint128 {
	lo, carry := bits.Add64(u.lo, m, 0)
	return uint128{u.hi + carry, lo}
}

// shiftLeft shifts u left by b
func (u uint128) shiftLeft(b uint64) uint128 {
	return uint128{
		(u.hi << b) | (u.lo >> (64 - b)) | (u.lo << (b - 64)),
		u.lo << b,
	}
}

// shiftLeft shifts u right by b
func (u uint128) shiftRight(b uint64) uint128 {
	return uint128{
		u.hi >> b,
		(u.lo >> b) | (u.hi << (64 - b)) | (u.hi >> (b - 64)),
	}
}

// bit returns the bit at position b
func (u uint128) bit(b uint64) uint8 {
	return uint8(u.shiftRight(b).lo & 1)
}

func (u uint128) String() string {
	f := "%0" + fmt.Sprint(64) + "b"
	s := []byte(fmt.Sprintf(f+f, u.hi, u.lo))
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
	return string(s)
}
