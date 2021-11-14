package p0425wordsquares

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordSquares(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  [][]string
	}{
		{[]string{"abat", "baba", "atan", "atal"}, [][]string{{"baba", "abat", "baba", "atal"}, {"baba", "abat", "baba", "atan"}}},
		{[]string{"area", "lead", "wall", "lady", "ball"}, [][]string{{"ball", "area", "lead", "lady"}, {"wall", "area", "lead", "lady"}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			res := wordSquares(tc.words)
			require.ElementsMatch(t, tc.want, res)
		})
	}
}

type trieNode struct {
	next  [26]*trieNode
	words []string
}

func (n *trieNode) getNext(ch byte) *trieNode {
	if n.next[ch] == nil {
		n.next[ch] = new(trieNode)
	}
	return n.next[ch]
}

func wordSquares(words []string) [][]string {
	// Add all words to a trie
	root := new(trieNode)
	for _, word := range words {
		cur := root
		cur.words = append(cur.words, word)
		for i := range word {
			cur = cur.getNext(word[i] - 'a')
			cur.words = append(cur.words, word)
		}
	}

	var f squareFinder
	f.root = root
	curSquare := make([]string, len(words[0]))
	f.findWordSquares(curSquare, 0, len(words[0]))
	return f.wordSquares
}

type squareFinder struct {
	root        *trieNode
	wordSquares [][]string
}

func (f *squareFinder) findWordSquares(curSquare []string, pos, n int) {
	if pos == n {
		stringsCpy := make([]string, len(curSquare))
		copy(stringsCpy, curSquare)
		f.wordSquares = append(f.wordSquares, stringsCpy)
		return
	}
	cur := f.root
	for i := 0; i < pos; i++ {
		cur = cur.next[curSquare[i][pos]-'a']
		if cur == nil {
			return
		}
	}
	for _, word := range cur.words {
		curSquare[pos] = word
		f.findWordSquares(curSquare, pos+1, n)
	}
}
