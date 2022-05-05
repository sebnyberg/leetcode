package p0127wordladder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLadders(t *testing.T) {
	for _, tc := range []struct {
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{"a", "c", []string{"a", "b", "c"}, 2},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0},
	} {
		t.Run(fmt.Sprintf("%v/%v/%+v", tc.beginWord, tc.endWord, tc.wordList), func(t *testing.T) {
			got := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
			require.Equal(t, tc.want, got)
		})
	}
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Find index of begin / end word
	var beginIdx, endIdx int = -1, -1
	for i, w := range wordList {
		if w == beginWord {
			beginIdx = i
		} else if w == endWord {
			endIdx = i
		}
	}
	if endIdx == -1 {
		return 0
	}

	// If beginWord is not in the list, add it
	if beginIdx == -1 {
		beginIdx = len(wordList)
		wordList = append(wordList, beginWord)
	}
	n := len(wordList)

	// Create wildcard masks for each word in the list, where one character is
	// replaced with '*'. Store these masks and indices of words having the masks.
	maskIndices := make(map[string][]int, n)
	masks := make([][]string, n)
	for i, w := range wordList {
		masks[i] = make([]string, len(beginWord))
		for j := 0; j < len(beginWord); j++ {
			m := string(w[:j] + "*" + w[j+1:])
			maskIndices[m] = append(maskIndices[m], i)
			masks[i][j] = m
		}
	}

	// Perform BFS until finding the endWord or running out of options
	seen := make([]bool, n)
	seen[beginIdx] = true
	cur := make([]int, 1, n)
	cur[0] = beginIdx
	next := make([]int, 0, n)
	for nwords := 2; len(cur) > 0; nwords++ {
		next = next[:0]
		for _, idx := range cur {
			for _, m := range masks[idx] {
				for _, nei := range maskIndices[m] {
					if seen[nei] {
						continue
					}
					if nei == endIdx {
						return nwords
					}
					seen[nei] = true
					next = append(next, nei)
				}
			}
		}

		cur, next = next, cur
	}

	return 0
}
