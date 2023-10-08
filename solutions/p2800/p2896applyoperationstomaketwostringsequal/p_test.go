package p2896applyoperationstomaketwostringsequal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for i, tc := range []struct {
		s1   string
		s2   string
		x    int
		want int
	}{
		{"1100011000", "0101001010", 2, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.s1, tc.s2, tc.x))
		})
	}
}

func minOperations(s1 string, s2 string, x int) int {
	// For a solution to exist, there must be an even mismatch of numbers
	var mismatch int
	for i := range s1 {
		if s1[i] != s2[i] {
			mismatch++
		}
	}
	if mismatch%2 != 0 {
		return -1
	}

	m := make(map[state]int)
	res := dp(m, s1, s2, s1[0], 0, x, false)
	return res
}

func flip(c byte) byte {
	a := c - '0'
	b := 1 - a
	return b + '0'
}

func dp(mem map[state]int, s1, s2 string, c byte, i, x int, hasXFlip bool) int {
	// For any position, we may either:
	//
	// 1) do nothing
	// 2) flip current and next
	// 3) flip current and some previous that has already been discounted
	// 4) flip current and some further number in the future
	//
	if i == len(s1)-1 {
		if hasXFlip {
			return 0
		}
		if s2[i] != c {
			panic("last character did not match, impossible case")
		}
		return 0
	}
	k := state{hasXFlip, c, i}
	if v, exists := mem[k]; exists {
		return v
	}

	// First try to do nothing
	if s2[i] == c {
		mem[k] = dp(mem, s1, s2, s1[i+1], i+1, x, hasXFlip)
		return mem[k]
	}

	// We must flip
	// Operation 1
	res := 1 + dp(mem, s1, s2, flip(s1[i+1]), i+1, x, hasXFlip)

	// Operation 2
	if hasXFlip {
		res = min(res, dp(mem, s1, s2, s1[i+1], i+1, x, false))
	} else {
		res = min(res, x+dp(mem, s1, s2, s1[i+1], i+1, x, true))
	}

	mem[k] = res
	return mem[k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type state struct {
	hasXflip bool
	currval  byte
	i        int
}
