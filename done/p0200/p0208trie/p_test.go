package p0208trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrieNode(t *testing.T) {
	trie := Constructor()
	for _, action := range []struct {
		name string
		arg  string
		want bool
	}{
		{"insert", "apple", false},
		{"search", "apple", true},
		{"search", "app", false},
		{"startsWith", "app", true},
		{"insert", "app", false},
		{"search", "app", true},
	} {
		switch action.name {
		case "insert":
			trie.Insert(action.arg)
		case "search":
			res := trie.Search(action.arg)
			require.Equal(t, action.want, res)
		case "startsWith":
			res := trie.StartsWith(action.arg)
			require.Equal(t, action.want, res)
		}
	}
}

type TrieNode struct {
	next map[rune]*TrieNode
	end  bool
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: new(TrieNode),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for _, ch := range word {
		if cur.next == nil {
			cur.next = make(map[rune]*TrieNode)
		}
		if cur.next[ch] == nil {
			cur.next[ch] = new(TrieNode)
		}
		cur = cur.next[ch]
	}
	cur.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, ch := range word {
		if _, exists := cur.next[ch]; !exists {
			return false
		}
		cur = cur.next[ch]
	}
	return cur.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, ch := range prefix {
		if _, exists := cur.next[ch]; !exists {
			return false
		}
		cur = cur.next[ch]
	}
	return true
}
