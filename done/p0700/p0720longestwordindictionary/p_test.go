package p0720longestwordindictionary

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestWord(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  string
	}{
		{[]string{"ogz", "eyj", "e", "ey", "hmn", "v", "hm", "ogznkb", "ogzn", "hmnm", "eyjuo", "vuq", "ogznk", "og", "eyjuoi", "d"}, "eyj"},
		{[]string{"w", "wo", "wor", "worl", "world"}, "world"},
		{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, "apple"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, longestWord(tc.words))
		})
	}
}

type trieNode struct {
	next  [26]*trieNode
	count int
	end   bool
	s     string
}

func longestWord(words []string) string {
	// Add all words to a trie, counting words on each node
	root := &trieNode{}
	root.end = true
	root.count = len(words) + 1
	for _, w := range words {
		curr := root
		for _, ch := range w {
			k := int(ch - 'a')
			if curr.next[k] == nil {
				curr.next[k] = &trieNode{}
			}
			curr = curr.next[k]
			curr.count++
		}
		curr.end = true
		curr.s = w
	}

	// Explore all paths in trie, storing max result
	var maxRes string
	dfs(root, math.MaxInt32, &maxRes)
	return maxRes
}

func dfs(curr *trieNode, k int, maxRes *string) {
	if !curr.end {
		return
	}
	if len(curr.s) > len(*maxRes) || curr.s < *maxRes {
		*maxRes = curr.s
	}
	k = min(k, curr.count)
	if k == 1 {
		return
	}
	for _, node := range curr.next {
		if node == nil {
			continue
		}
		dfs(node, k-1, maxRes)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
