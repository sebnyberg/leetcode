package p0336palindromepairs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_palindromePairs(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  [][]int
	}{
		{[]string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
		{[]string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}}},
		{[]string{"a", ""}, [][]int{{0, 1}, {1, 0}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, palindromePairs(tc.words))
		})
	}
}

func palindromePairs(words []string) [][]int {
	// Combining and checking all pairs of words is O((n*m)^2),
	// where:
	// n is the number of words, and
	// m is the length of the combined word.
	//
	// However, when considering a certain word, there are only
	// so many ways to combine that word with another word to
	// form a new word.
	// For example, we start with sorting words by length.
	// Then for each word, we check all possible words that the
	// word could be combined with to form a palindrome, such that
	// the other word has shorter or equal length to the current
	// word. Any such word is counted.
	//
	// First, we create an auxiliary item which retains the original
	// position along with the word
	//
	// Then, let's consider the example words:
	// ["a", "lls", "abcd", "dcba", "sssll"]
	//
	// We need to find all candidates which could be added to "a" on
	// either its left or right side such that the candidate is smaller
	// or equal in length to "a".
	//
	// We can do that by moving a window on left and right side such that the
	// windows left or right boundary always falls within the word, and the other
	// boundary falls outside.
	//
	// If the missing portion of the word has been seen already, then there is a
	// match.
	buf := []byte{}
	findMatches := func(seenAtIdx map[string][]int, l, r int, word string) []int {
		buf = buf[:0]
		doReverse := r >= len(word)
		for ; l < r; l, r = l+1, r-1 {
			if l < 0 {
				buf = append(buf, word[r])
			} else if r >= len(word) {
				buf = append(buf, word[l])
			} else {
				if word[l] != word[r] {
					return nil
				}
			}
		}
		if doReverse {
			for ll, rr := 0, len(buf)-1; ll < rr; ll, rr = ll+1, rr-1 {
				buf[ll], buf[rr] = buf[rr], buf[ll]
			}
		}
		return seenAtIdx[string(buf)]
	}
	type item struct {
		idx  int
		word string
	}
	n := len(words)
	ws := make([]item, n)
	for i, w := range words {
		ws[i] = item{idx: i, word: w}
	}
	sort.Slice(ws, func(i, j int) bool {
		return len(ws[i].word) < len(ws[j].word)
	})
	seenAtIdx := make(map[string][]int)
	isPalindrome := func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}

	var res [][]int
	for _, w := range ws {
		// Move a window over the word and find possible matches
		for l := -len(w.word); l < 0; l++ {
			for _, w2Idx := range findMatches(seenAtIdx, l, len(w.word)-1, w.word) {
				res = append(res, []int{w2Idx, w.idx})
			}
		}
		// Move a window over the word and find possible matches
		for r := len(w.word)*2 - 1; r >= len(w.word); r-- {
			for _, w2Idx := range findMatches(seenAtIdx, 0, r, w.word) {
				res = append(res, []int{w.idx, w2Idx})
			}
		}
		// Annoying, special case of adding nothing at all...
		if isPalindrome(w.word) {
			for _, w2Idx := range seenAtIdx[""] {
				res = append(res, []int{w2Idx, w.idx}, []int{w.idx, w2Idx})
			}
		}
		seenAtIdx[w.word] = append(seenAtIdx[w.word], w.idx)
	}
	return res
}
