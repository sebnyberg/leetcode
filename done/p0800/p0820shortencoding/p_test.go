package p0820shortencoding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumLengthEncoding(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  int
	}{
		{[]string{"time", "atime", "btime"}, 12},
		{[]string{"time", "me", "bell"}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, minimumLengthEncoding(tc.words))
		})
	}
}

type TrieNode struct {
	val      rune
	children map[rune]*TrieNode
}

func (n *TrieNode) GetOrPut(r rune) *TrieNode {
	if _, exists := n.children[r]; !exists {
		n.children[r] = &TrieNode{
			val:      r,
			children: make(map[rune]*TrieNode),
		}
	}
	return n.children[r]
}

func minimumLengthEncoding(words []string) int {
	// Create a trie with the words from end to start
	root := &TrieNode{
		val:      rune(0),
		children: make(map[rune]*TrieNode),
	}
	for _, word := range words {
		cur := root
		n := len(word)
		for i := range word {
			cur = cur.GetOrPut(rune(word[n-i-1]))
		}
	}

	// Once the trie is completed, the return value is the sum of
	// the depth of all paths in the Trie * number of paths (#)
	npaths, sum := findPaths(root, 0)

	return npaths + sum
}

func findPaths(cur *TrieNode, cursum int) (npaths int, sum int) {
	if len(cur.children) == 0 {
		return 1, cursum
	}
	for _, path := range cur.children {
		nsubpaths, subsum := findPaths(path, cursum+1)
		npaths += nsubpaths
		sum += subsum
	}
	return npaths, sum
}
