package p0211designaddsearchwordsdatastructure

import (
	"fmt"
	"testing"

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

type WordDictionary struct {
	root *node
	curr []*node
	next []*node
}

type node struct {
	next [26]*node
	end  bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &node{},
		curr: []*node{},
		next: []*node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, ch := range word {
		ch -= 'a'
		if curr.next[ch] == nil {
			curr.next[ch] = &node{}
		}
		curr = curr.next[ch]
	}
	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	// Using BFS here, trading code complexity for better performance
	this.curr = append(this.curr[:0], this.root)
	for len(this.curr) > 0 {
		this.next = this.next[:0]
		var end bool
		ch := word[0] - 'a'
		for _, x := range this.curr {
			if ch+'a' == '.' {
				// add all non-nil edges
				for _, y := range x.next {
					if y != nil {
						this.next = append(this.next, y)
						end = end || y.end
					}
				}
			} else {
				// only add matching edges
				y := x.next[ch]
				if y != nil {
					this.next = append(this.next, y)
					end = end || y.end
				}
			}
		}
		word = word[1:]
		if len(word) == 0 {
			return end
		}
		this.curr, this.next = this.next, this.curr
	}
	return false
}
