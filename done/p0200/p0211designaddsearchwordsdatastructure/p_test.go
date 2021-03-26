package p0211designaddsearchwordsdatastructure

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func TestWordDictionary(t *testing.T) {
	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	for _, tc := range []struct {
		in   string
		want bool
	}{
		{"pad", false},
		{"bad", true},
		{".ad", true},
		{"b..", true},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, dict.Search(tc.in))
		})
	}
}

type TrieNode struct {
	next map[rune]*TrieNode
	end  bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	root := &TrieNode{
		next: make(map[rune]*TrieNode),
	}
	return WordDictionary{root: root}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, ch := range word {
		if _, exists := cur.next[ch]; !exists {
			cur.next[ch] = &TrieNode{
				next: make(map[rune]*TrieNode),
			}
		}
		cur = cur.next[ch]
	}
	cur.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchNode(this.root, word)
}

func (this *WordDictionary) searchNode(cur *TrieNode, word string) bool {
	if len(word) == 0 && cur.end {
		return true
	}
	ch, width := utf8.DecodeRuneInString(word)
	if ch != '.' {
		if next, exists := cur.next[ch]; exists {
			return this.searchNode(next, word[width:])
		}
		return false
	}
	// ch == '.'
	for _, next := range cur.next {
		if this.searchNode(next, word[width:]) {
			return true
		}
	}
	return false
}
