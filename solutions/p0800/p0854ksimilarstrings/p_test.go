package p0854ksimilarstrings

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kSimilarity(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		s2   string
		want int
	}{
		{"bdeabcefabcd", "ccadeadbbefb", 8},
		{"abcdefabcdefabcdef", "acccafdeaddbbefbef", 8},
		{"aabccb", "bbcaca", 3},
		{"ab", "ba", 1},
		{"abc", "bca", 2},
		{"abac", "baca", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s1), func(t *testing.T) {
			require.Equal(t, tc.want, kSimilarity(tc.s1, tc.s2))
		})
	}
}

func Test_newBM(t *testing.T) {
	a := "abc"
	b := "cba"
	bma := newBM(a)
	bmb := newBM(b)
	bma = swap(bma, 0, 2)
	require.Equal(t, bma, bmb)
}

func newBM(s string) uint64 {
	var res uint64
	for i := range s {
		res <<= 3
		res += uint64(s[i] - 'a' + 1)
	}
	return res
}

func swap(bm uint64, i, j int) uint64 {
	xored := val(bm, i) ^ val(bm, j)
	mask := xored<<(i*3) | xored<<(j*3)
	bm ^= mask
	return bm
}

func val(bm uint64, i int) uint64 {
	res := (bm >> (i * 3)) & 7
	return res
}

func kSimilarity(s1 string, s2 string) int {
	mem := make(map[uint64]int)
	res := fix(mem, newBM(s1), newBM(s2))
	return res
}

func fix(mem map[uint64]int, bm1, bm2 uint64) int {
	if bm1 == bm2 {
		return 0
	}
	if v, exists := mem[bm1]; exists {
		return v
	}
	if val(bm1, 0) == val(bm2, 0) {
		return fix(mem, bm1>>3, bm2>>3)
	}
	// Must swap with one of options in s1
	minCost := math.MaxInt32
	want := val(bm2, 0)
	for j := 1; val(bm1, j) > 0; j++ {
		if val(bm1, j) != want {
			continue
		}
		// Try swapping this position
		cpy := swap(bm1, 0, j)
		subRes := 1 + fix(mem, cpy>>3, bm2>>3)
		minCost = min(minCost, subRes)
	}
	mem[bm1] = minCost
	return minCost

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
