package p0212wordsearch2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWords(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{[][]byte{{'a'}}, []string{"a"}, []string{"a"}},
		{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abd", "acd"}, []string{"abd", "acd"}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.board, tc.words), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findWords(tc.board, tc.words))
		})
	}
}

func findWords(board [][]byte, words []string) []string {
	// create trie
	root := &trieNode{
		children: make(map[byte]*trieNode),
	}
	for _, word := range words {
		node := root
		for i := range word {
			ch := word[i]
			if _, exists := node.children[ch]; !exists {
				node.children[ch] = &trieNode{
					children: make(map[byte]*trieNode),
				}
			}
			node = node.children[ch]
		}
		node.val = word
	}

	// Keep track of the positions that have been visited
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range board {
		visited[i] = make([]bool, n)
	}

	// For each position on the board
	f := wordFinder{
		words:   make(map[string]struct{}),
		board:   board,
		visited: visited,
		n:       n,
		m:       m,
	}
	for i := range board {
		for j := range board[i] {
			f.findWordsWithTrie(root, i, j)
		}
	}

	result := make([]string, 0, len(f.words))
	for word := range f.words {
		result = append(result, word)
	}

	return result
}

type wordFinder struct {
	words   map[string]struct{}
	visited [][]bool
	board   [][]byte
	n       int
	m       int
}

func (f *wordFinder) invalidPosition(i, j int) bool {
	return i >= f.m || i < 0 || j >= f.n || j < 0
}

func (f *wordFinder) findWordsWithTrie(trie *trieNode, i, j int) {
	if f.invalidPosition(i, j) || f.visited[i][j] {
		return
	}
	if _, exists := trie.children[f.board[i][j]]; !exists {
		return
	}
	f.visited[i][j] = true
	trie = trie.children[f.board[i][j]]
	if trie.val != "" {
		f.words[trie.val] = struct{}{}
	}
	// visit adjacent positions
	for _, pos := range [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
		f.findWordsWithTrie(trie, pos[0], pos[1])
	}
	f.visited[i][j] = false
}

type trieNode struct {
	children map[byte]*trieNode
	val      string // Terminal nodes are marked with a value
}
