package p0677mapsumpairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapSum(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	res := mapSum.Sum("ap")
	require.Equal(t, 3, res)
	mapSum.Insert("app", 2)
	mapSum.Insert("apple", 2)
	res = mapSum.Sum("ap")
	require.Equal(t, 4, res)
}

type MapSum struct {
	values map[string]int
	trie   *TrieNode
}

type TrieNode struct {
	next [26]*TrieNode
	val  int
}

func (n *TrieNode) Next(ch rune) *TrieNode {
	if n.next[ch-'a'] == nil {
		n.next[ch-'a'] = &TrieNode{}
	}
	return n.next[ch-'a']
}

func Constructor() MapSum {
	return MapSum{
		values: make(map[string]int),
		trie:   &TrieNode{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	if oldVal, exists := this.values[key]; exists {
		delta -= oldVal
	}
	this.values[key] = val

	// Add to trie
	cur := this.trie
	for _, ch := range key {
		cur = cur.Next(ch)
		cur.val += delta
	}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.trie
	for _, ch := range prefix {
		cur = cur.Next(ch)
	}
	return cur.val
}
