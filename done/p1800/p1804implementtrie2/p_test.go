package p1804implementtrie2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrie2(t *testing.T) {
	tr := Constructor()
	tr.Insert("n")
	tr.Insert("n")
	tr.Insert("jvo")
	tr.Erase("n")
	res := tr.CountWordsStartingWith("n")
	require.Equal(t, 1, res)
	res = tr.CountWordsEqualTo("n")
	require.Equal(t, 1, res)
	tr.Erase("n")
	res = tr.CountWordsStartingWith("n")
	require.Equal(t, 0, res)
	res = tr.CountWordsEqualTo("n")
	require.Equal(t, 0, res)
	tr.Insert("n")
}

type TrieNode struct {
	children [26]*TrieNode
	starts   int // words on this path
	ends     int // words ending here
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := range word {
		cur.starts++
		a := word[i] - 'a'
		if cur.children[a] == nil {
			cur.children[a] = &TrieNode{}
		}
		cur = cur.children[a]
	}
	cur.ends++
}

func (this *Trie) CountWordsEqualTo(word string) int {
	cur := this.root
	for i := range word {
		a := word[i] - 'a'
		if cur.children[a] == nil {
			return 0
		}
		cur = cur.children[a]
	}
	return cur.ends
}

func (this *Trie) CountWordsStartingWith(prefix string) int {
	cur := this.root
	for i := range prefix {
		a := prefix[i] - 'a'
		if cur.children[a] == nil {
			return 0
		}
		cur = cur.children[a]
	}
	return cur.starts + cur.ends
}

func (this *Trie) Erase(word string) {
	cur := this.root
	for i := range word {
		cur.starts--
		a := word[i] - 'a'
		if cur.children[a] == nil {
			return
		}
		cur = cur.children[a]
	}
	cur.ends--
}
