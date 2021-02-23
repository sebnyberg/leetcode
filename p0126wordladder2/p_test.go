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
	considered := make([]bool, len(wordList))
	toConsider := make([]int, 0)
	lastWordIdx := -1
	for i, word := range wordList {
		if word == endWord {
			lastWordIdx = i
		}
		if isAdj(beginWord, word) {
			toConsider = append(toConsider, i)
			if word == endWord {
				return 2
			}
			considered[i] = true
		}
	}
	if lastWordIdx == -1 {
		return 0
	}

	listLen := 1
	for len(toConsider) > 0 {
		newIndices := make([]int, 0, 10)
		for i, otherWord := range wordList {
			if considered[i] {
				continue
			}
			for _, considerIdx := range toConsider {
				if isAdj(wordList[considerIdx], otherWord) {
					if i == lastWordIdx {
						return listLen + 2
					}
					considered[i] = true
					newIndices = append(newIndices, i)
				}
			}
		}
		toConsider = newIndices
		listLen++
	}

	return 0
}

func isAdj(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var dist int
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
		if dist > 1 {
			return false
		}
	}
	return dist == 1
}
