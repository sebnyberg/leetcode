package p0642designsearchautocompletesystem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type any interface{}

func TestAutocomplete(t *testing.T) {
	type action struct {
		name string
		args []byte
		want []any
	}
	for i, tc := range []struct {
		inp     []string
		inp2    []int
		actions []action
	}{
		{
			[]string{"i love you", "island", "iroman", "i love leetcode"},
			[]int{5, 3, 2, 2},
			[]action{
				{"input", []byte{'i'}, []any{[]string{"i love you", "island", "i love leetcode"}}},
				{"input", []byte{' '}, []any{[]string{"i love you", "i love leetcode"}}},
				{"input", []byte{'a'}, []any{[]string{}}},
				{"input", []byte{'#'}, []any{[]string{}}},
			},
		},
		{
			[]string{"i love you", "island", "iroman", "i love leetcode"},
			[]int{5, 3, 2, 2},
			[]action{
				{"input", []byte{'i'}, []any{[]string{"i love you", "island", "i love leetcode"}}},
				{"input", []byte{' '}, []any{[]string{"i love you", "i love leetcode"}}},
				{"input", []byte{'a'}, []any{[]string{}}},
				{"input", []byte{'#'}, []any{[]string{}}},
				{"input", []byte{'i'}, []any{[]string{"i love you", "island", "i love leetcode"}}},
				{"input", []byte{' '}, []any{[]string{"i love you", "i love leetcode", "i a"}}},
				{"input", []byte{'a'}, []any{[]string{"i a"}}},
				{"input", []byte{'#'}, []any{[]string{}}},
				{"input", []byte{'i'}, []any{[]string{"i love you", "island", "i a"}}},
				{"input", []byte{' '}, []any{[]string{"i love you", "i a", "i love leetcode"}}},
				{"input", []byte{'a'}, []any{[]string{"i a"}}},
				{"input", []byte{'#'}, []any{[]string{}}},
			},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := Constructor(tc.inp, tc.inp2)
			for j, act := range tc.actions {
				res := c.Input(act.args[0])
				require.ElementsMatch(t, act.want[0], res, j)
			}
		})
	}
}

type AutocompleteSystem struct {
	stack []*trieNode
	root  *trieNode
	input []byte
}

type trieNode struct {
	next      [27]*trieNode
	sentences []sentence
}

func (n *trieNode) getNext(i int) *trieNode {
	if i < 0 { // ' '
		i = 26
	}
	if n.next[i] == nil {
		n.next[i] = new(trieNode)
	}
	return n.next[i]
}

func (n *trieNode) add(s string, x int) {
	for i, t := range n.sentences {
		if t.s == s {
			n.sentences[i].count += x
			return
		}
	}
	n.sentences = append(n.sentences, sentence{s: s, count: x})
}

type sentence struct {
	count int
	s     string
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	a := AutocompleteSystem{
		root: new(trieNode),
	}
	for i, s := range sentences {
		cur := a.root
		for _, ch := range s {
			cur = cur.getNext(int(ch) - 'a')
			cur.add(s, times[i])
		}
	}
	return a
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		// Add the sentence to the system
		s := string(this.input)
		for i := 0; i < len(this.stack); i++ {
			this.stack[i].add(s, 1)
		}
		this.stack = this.stack[:0]
		this.input = this.input[:0]
		return []string{}
	}
	this.input = append(this.input, c)
	var cur *trieNode
	if len(this.stack) == 0 {
		cur = this.root.getNext(int(c) - 'a')
	} else {
		cur = this.stack[len(this.stack)-1].getNext(int(c) - 'a')
	}
	this.stack = append(this.stack, cur)
	sort.Slice(cur.sentences, func(i, j int) bool {
		if cur.sentences[i].count == cur.sentences[j].count {
			return cur.sentences[i].s < cur.sentences[j].s
		}
		return cur.sentences[i].count > cur.sentences[j].count
	})
	var res []string
	for i := 0; i < min(3, len(cur.sentences)); i++ {
		res = append(res, cur.sentences[i].s)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
