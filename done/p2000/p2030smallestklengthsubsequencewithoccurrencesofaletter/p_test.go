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
	// Pick as many as you can from the string so that the prior letter is
	return ""
}
