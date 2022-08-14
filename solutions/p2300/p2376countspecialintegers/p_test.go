package p2376countspecialintegers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSpecialNumbers(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{500, 378},
		{225, 178},
		{135, 110},
		{5, 5},
		{20, 19},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countSpecialNumbers(tc.n))
		})
	}
}

// https://www.github.com/sebnyberg/leetcode
func countSpecialNumbers(n int) int {
	// There are 2^9 different sets of digits, which is quite small.
	// Then there are up to 9 different lengths.
	// This means we can memoize the number of distinct integers of a certain
	// length given a set of already used integers.
	var mem [1024][10]int
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	x := fmt.Sprint(n)

	// Count all special numbers with fewer digits than  n
	var res int
	for k := len(x) - 1; k >= 1; k-- {
		res += countValid(&mem, 0, k)
	}

	// Count all special numbers smaller than some set values
	var bm int
	for i := 0; i < len(x); i++ {
		upper := int(x[i] - '0')
		for y := 0; y < upper; y++ {
			if bm&(1<<y) > 0 {
				continue
			}
			res += countValid(&mem, bm|(1<<y), len(x)-i-1)
		}
		if bm&(1<<upper) > 0 {
			return res
		}
		bm |= (1 << upper)
	}
	return res + 1
}

func countValid(mem *[1024][10]int, bm, nleft int) int {
	if bm == 1 {
		return 0
	}
	if nleft == 0 {
		return 1
	}
	if mem[bm][nleft] != -1 {
		return mem[bm][nleft]
	}
	var res int
	for x := 0; x <= 9; x++ {
		if bm&(1<<x) == 0 {
			res += countValid(mem, bm|(1<<x), nleft-1)
		}
	}
	mem[bm][nleft] = res
	return res
}
