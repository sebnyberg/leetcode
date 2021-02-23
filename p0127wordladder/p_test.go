package p0126wordladder2

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
		want      [][]string
	}{
		{"a", "c", []string{"a", "b", "c"}, [][]string{{"a", "c"}}},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, [][]string{
			{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"},
		}},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, [][]string{}},
	} {
		t.Run(fmt.Sprintf("%v/%v/%+v", tc.beginWord, tc.endWord, tc.wordList), func(t *testing.T) {
			got := findLadders(tc.beginWord, tc.endWord, tc.wordList)
			require.Equal(t, tc.want, got)
		})
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	considered := make([]bool, len(wordList))
	resIndices := make([][]int, 0)
	lastWordIdx := -1
	for i, word := range wordList {
		if word == endWord {
			lastWordIdx = i
		}
		if isAdj(beginWord, word) {
			resIndices = append(resIndices, []int{i})
			considered[i] = true
		}
	}
	if lastWordIdx == -1 {
		return [][]string{}
	}

	listLen := 1
	for len(resIndices) > 0 {
		if considered[lastWordIdx] {
			break
		}
		newIndices := make([][]int, 0, 10)
		for i, otherWord := range wordList {
			if considered[i] {
				continue
			}
			for _, words := range resIndices {
				if isAdj(wordList[words[listLen-1]], otherWord) {
					considered[i] = true
					wordsCpy := make([]int, listLen+1)
					copy(wordsCpy, words)
					wordsCpy[listLen] = i
					newIndices = append(newIndices, wordsCpy)
				}
			}
		}
		resIndices = newIndices
		listLen++
	}

	res := make([][]string, 0, len(resIndices))
	for _, r := range resIndices {
		if r[listLen-1] != lastWordIdx {
			continue
		}
		ws := make([]string, listLen+1)
		ws[0] = beginWord
		for j, idx := range r {
			ws[j+1] = wordList[idx]
		}
		res = append(res, ws)
	}
	return res
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
