package p2030smallestklengthsubsequencewithoccurrencesofaletter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s          string
		k          int
		letter     byte
		repetition int
		want       string
	}{
		{"facfffkfnffoppfffzfz", 9, 'f', 9, "fffffffff"},
		{"adffhjfmmmmorsfff", 6, 'f', 5, "afffff"},
		{"aaabbbcccddd", 3, 'b', 2, "abb"},
		{"leet", 3, 'e', 1, "eet"},
		{"leetcode", 4, 'e', 2, "ecde"},
		{"bb", 2, 'b', 2, "bb"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, smallestSubsequence(tc.s, tc.k, tc.letter, tc.repetition))
		})
	}
}

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	// Count occurrences of letter in s so that it's know how many can be
	// discarded.
	var letterCount int
	for i := range s {
		if s[i] == letter {
			letterCount++
		}
	}
	discardable := letterCount - repetition

	// Keep a kind of monotonic stack of letters. Since it's constrained by the
	// number of available instances of 'letter', it will not be truly monotonic.
	stack := make([]byte, 0)
	for i := range s {
		// Pop from the stack until either:
		// * The new letter would not improve the lexicographical ordering, or
		// * There are not enough characters left to form a k-width result, or
		// * We'd pop a 'letter' such that we could not meet 'repetition' instances.
		for len(stack) != 0 &&
			len(s)-i+len(stack) > k &&
			stack[len(stack)-1] > s[i] &&
			(stack[len(stack)-1] != letter || discardable > 0) {
			if stack[len(stack)-1] == letter {
				discardable--
				repetition++
			}
			stack = stack[:len(stack)-1]
		}

		if len(stack) < k {
			if s[i] == letter {
				repetition--
				stack = append(stack, s[i])
			} else if k-len(stack) > repetition {
				stack = append(stack, s[i])
			}
		} else if s[i] == letter {
			discardable--
		}
	}

	return string(stack)
}
