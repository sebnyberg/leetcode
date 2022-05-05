package p2135countwordsobtainedafteraddingaletter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordCount(t *testing.T) {
	for _, tc := range []struct {
		startWords  []string
		targetWords []string
		want        int
	}{
		{[]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}, 2},
		{[]string{"ab", "a"}, []string{"abc", "abcd"}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.startWords), func(t *testing.T) {
			require.Equal(t, tc.want, wordCount(tc.startWords, tc.targetWords))
		})
	}
}

func wordCount(startWords []string, targetWords []string) int {
	// Since characters can be rearranged, it does not matter which order they
	// are in - counts matters more.
	// Also, since no existing character can be added to the string, the count
	// for characters that exist in the target word must exactly match the
	// startWord.

	// So, in short, a targetWord can be formed from a start word IF a start word
	// exists such that its counts of characters exactly correspond to the counts
	// of characters found in targetWord (not vice versa necessarily)

	// Just found the constraint about max one letter per word... jesus

	// If there is any subset of picked letters from targetWords which exists
	// in startWords, then there is a match.
	bmFromWord := func(w string) int {
		var bm int
		for i := range w {
			bm |= 1 << int(w[i]-'a')
		}
		return bm
	}

	start := make(map[int]struct{}, len(startWords))
	for _, w := range startWords {
		start[bmFromWord(w)] = struct{}{}
	}

	// For each target word, if a startword exists such that a set (or subset) of
	// characters in target word exists, then there is a match.
	var res int
	for _, w := range targetWords {
		bm := bmFromWord(w)
		for i := 0; i < len(w); i++ {
			if _, exists := start[bm&^(1<<int(w[i]-'a'))]; exists {
				res++
				break
			}
		}
	}
	return res
}
