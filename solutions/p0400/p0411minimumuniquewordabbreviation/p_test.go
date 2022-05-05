package p0411minimumuniquewordabbreviation

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minAbbreviation(t *testing.T) {
	for _, tc := range []struct {
		target     string
		dictionary []string
		want       string
	}{
		{"apple", []string{"kkkk", "blade"}, "a4"},
		{"apple", []string{"blade"}, "a4"},
		{"apple", []string{"blade", "plain", "amber"}, "1p3"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, minAbbreviation(tc.target, tc.dictionary))
		})
	}
}

func minAbbreviation(target string, dictionary []string) string {
	n := len(target)

	// Helpers
	wordBits := func(w string) int {
		var bitWord int
		for i := range w {
			bitWord <<= 1
			if target[i] != w[i] {
				bitWord |= 1
			}
		}
		return bitWord
	}
	abbrLen := func(mask int) int {
		count := n
		for b := 3; b < (1 << n); b <<= 1 {
			if mask&b == 0 {
				count--
			}
		}
		return count
	}

	// Pre-process dictionary words to contain 1s in the positions where the
	// target and dictionary characters are different. The solution must be a
	// bitmask such that it holds characters in these positions, i.e. we are
	// looking for a bitmask such that bitmask&dictionary[i] > 0 for all words
	// in the dictionary.
	var candBits int
	bitWords := make([]int, 0, len(dictionary))
	for _, w := range dictionary {
		if len(w) != n {
			continue
		}
		bits := wordBits(w)
		candBits |= bits
		bitWords = append(bitWords, bits)
	}

	// DFS places a number of bits in the mask based on where the candidate bits
	// lie. The match then (potentially) updates the minLength.
	minLength := math.MaxInt32
	var minResult int
	var dfs func(idx, bitsLeft, mask int)
	dfs = func(idx, bitsLeft, mask int) {
		maskLen := abbrLen(mask)
		if maskLen >= minLength {
			return
		}
		if idx == n || bitsLeft == 0 {
			if bitsLeft > 0 { // did not place enough bits
				return
			}
			// Check if the mask is a match
			for _, w := range bitWords {
				if mask&w == 0 {
					return
				}
			}
			// It is a match, maybe update minLength
			if maskLen < minLength {
				minResult = mask
				minLength = maskLen
			}
			return
		}

		// First, skip this position
		dfs(idx+1, bitsLeft, mask)

		// If this bit is a candidate bit, try to put a bit in this position
		if (1<<idx)&candBits == 0 {
			return
		}
		dfs(idx+1, bitsLeft-1, mask|(1<<idx))
	}

	// Perform DFS to find the minimal length abbreviation
	nbitsMax := bits.OnesCount(uint(candBits))
	for nbits := 0; nbits <= nbitsMax; nbits++ {
		dfs(0, nbits, 0)
	}

	// Convert mask to word
	res := make([]byte, 0, n)
	var zeroCount int
	for i := 0; i < n; i++ {
		bit := 1 << (n - i - 1)
		if minResult&bit > 0 {
			if zeroCount > 0 {
				res = append(res, []byte(strconv.Itoa(zeroCount))...)
			}
			zeroCount = 0
			res = append(res, target[i])
			continue
		}
		zeroCount++
	}
	if zeroCount > 0 {
		res = append(res, []byte(strconv.Itoa(zeroCount))...)
	}

	return string(res)
}
