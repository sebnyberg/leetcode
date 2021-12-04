package p1032streamofcharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StreamChecker(t *testing.T) {
	for _, tc := range []struct {
		words   []string
		queries []byte
		want    []bool
	}{
		{[]string{"cd", "f", "kl"}, []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}, []bool{false, false, false, true, false, true, false, false, false, false, false, true}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			checker := Constructor(tc.words)
			for i, query := range tc.queries {
				require.Equal(t, tc.want[i], checker.Query(query))
			}
		})
	}
}

type trieNode struct {
	next [26]*trieNode
	end  bool
}

func (n *trieNode) add(ch byte) *trieNode {
	if n.next[idx(ch)] == nil {
		n.next[idx(ch)] = new(trieNode)
	}
	return n.next[idx(ch)]
}

func idx(ch byte) byte {
	return ch - 'a'
}

type StreamChecker struct {
	root *trieNode
	word []byte
}

func Constructor(words []string) StreamChecker {
	// Construct a trie of input words
	sc := StreamChecker{}
	sc.root = new(trieNode)
	sc.word = make([]byte, 0, 1000)
	for _, w := range words {
		cur := sc.root
		for i := len(w) - 1; i >= 0; i-- {
			cur = cur.add(w[i])
		}
		cur.end = true
	}
	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.word = append(this.word, letter)
	cur := this.root
	for i := len(this.word) - 1; i >= 0; i-- {
		next := cur.next[idx(this.word[i])]
		if next == nil {
			return false
		}
		if next.end == true {
			return true
		}
		cur = next
	}
	return false
}
