package p2068checkwhethertwostringsarealmostequivalent

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkAlmostEquivalent(t *testing.T) {
	for _, tc := range []struct {
		word1 string
		word2 string
		want  bool
	}{
		{"aaaa", "bccb", false},
		{"abcdeef", "abaaacc", true},
		{"cccddabba", "babababab", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word1), func(t *testing.T) {
			require.Equal(t, tc.want, checkAlmostEquivalent(tc.word1, tc.word2))
		})
	}
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var charCount [26]int
	for _, ch := range word1 {
		charCount[int(ch-'a')]++
	}
	for _, ch := range word2 {
		charCount[int(ch-'a')]--
	}
	for _, count := range charCount {
		if abs(count) > 3 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
