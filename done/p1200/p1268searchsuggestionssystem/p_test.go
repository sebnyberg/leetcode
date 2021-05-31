package p1268searchsuggestionssystem

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_suggestedProducts(t *testing.T) {
	for _, tc := range []struct {
		products   []string
		searchWord string
		want       [][]string
	}{
		{[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse",
			[][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{[]string{"havana"}, "havana",
			[][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{[]string{"bags", "baggage", "banner", "box", "cloths"}, "bags",
			[][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"},
				{"bags"},
			},
		},
		{[]string{"havana"}, "tatiana",
			[][]string{{}, {}, {}, {}, {}, {}, {}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.products), func(t *testing.T) {
			require.Equal(t, tc.want, suggestedProducts(tc.products, tc.searchWord))
		})
	}
}

type trieNode struct {
	words []int
	next  [26]*trieNode
}

func (n *trieNode) Next(ch rune) *trieNode {
	if n.next[ch-'a'] == nil {
		n.next[ch-'a'] = &trieNode{
			words: make([]int, 0, 3),
		}
	}
	return n.next[ch-'a']
}

func suggestedProducts(products []string, searchWord string) [][]string {
	// Sort lexicographically
	sort.Strings(products)

	// Create a trie
	root := &trieNode{words: make([]int, 0)}

	for i, word := range products {
		cur := root
		for _, ch := range word {
			cur = cur.Next(ch)
			cur.words = append(cur.words, i)
		}
	}

	res := make([][]string, len(searchWord))
	cur := root
	for i, ch := range searchWord {
		cur = cur.Next(ch)
		res[i] = make([]string, 0, 3)
		for j := 0; j < min(3, len(cur.words)); j++ {
			res[i] = append(res[i], products[cur.words[j]])
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
