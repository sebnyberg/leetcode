package p0466counttherepetitions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMaxRepetitions(t *testing.T) {
	for _, tc := range []struct {
		s1   string
		n1   int
		s2   string
		n2   int
		want int
	}{
		{"acb", 4, "ab", 2, 2},
		{"acb", 1, "acb", 1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s1), func(t *testing.T) {
			require.Equal(t, tc.want, getMaxRepetitions(tc.s1, tc.n1, tc.s2, tc.n2))
		})
	}
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// The problem description is really bad for this problem, which is probably
	// why the acceptance rate is so low..
	//
	// If a characters can be removed from one string to match another string,
	// then there is a valid matching.
	//
	// The goal is to find a pattern of repeating s1s which amount to the same
	// amount of s2s every time it is repeated.
	//
	// For each instance of s1 that is being repeated, record the next index that
	// will be matched in s2. If the same index is requested for two distinct
	// instances of s1, then the distance between those two occurrences is the
	// number of repeats per period of s1s.
	//
	// Note that there may be a prefix prior to the start of the period, and a
	// suffix at the end. These need to be accounted for as well when determining
	// the max number of repeats of s2 in the answer.

	// nextRepeatIdx[j] = k means that the next index to match in s2 was j when
	// at repeat k
	nextRepeatIdx := make(map[int]int, len(s2))
	nextRepeatIdx[0] = 0 // must match index 0 in first repeat

	// repeats[k] = x means that after k repeats of s1, there were x repeats of s2
	repeats := make([]int, 1, len(s2))

	var nextIdx, count int
	for k := 1; k <= n1; k++ {
		// Count number of repeats of s2 in s1
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[nextIdx] {
				continue
			}
			nextIdx++
			if nextIdx == len(s2) {
				nextIdx = 0
				count++
			}
		}

		// Check if nextIdx has been seen before
		start, exists := nextRepeatIdx[nextIdx]
		if !exists { // Not seen, continue loop
			repeats = append(repeats, count)
			nextRepeatIdx[nextIdx] = k
			continue
		}

		repeatLen := n1 - start
		patternLen := k - start
		prefixCount := repeats[start]
		patternCount := (repeatLen / patternLen) * (count - prefixCount)
		suffixCount := repeats[start+repeatLen%patternLen] - prefixCount
		return (prefixCount + patternCount + suffixCount) / n2
	}
	return repeats[n1] / n2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
