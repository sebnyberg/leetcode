package p0212wordsearch2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWords(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abd", "acd"}, []string{"abd", "acd"}},
		{[][]byte{{'a'}}, []string{"a"}, []string{"a"}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.board, tc.words), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findWords(tc.board, tc.words))
		})
	}
}

type trieNode struct {
	next [26]*trieNode
	word string
}

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	n := len(board[0])

	root := &trieNode{}
	for _, w := range words {
		curr := root
		for _, ch := range w {
			if curr.next[ch-'a'] == nil {
				curr.next[ch-'a'] = &trieNode{}
			}
			curr = curr.next[ch-'a']
		}
		curr.word = w
	}

	var visited [12][12]bool
	seen := make(map[string]struct{})
	for i := range board {
		for j := range board[i] {
			explore(board, seen, root, &visited, i, j, m, n)
		}
	}
	var res []string
	for w := range seen {
		res = append(res, w)
	}
	return res
}

// explore the board from a given starting position, following routes in the
// trie and adding words to seen.
func explore(
	board [][]byte,
	seen map[string]struct{},
	curr *trieNode,
	visited *[12][12]bool,
	i, j, m, n int,
) {
	next := curr.next[board[i][j]-'a']
	if next == nil {
		return
	}
	if next.word != "" {
		seen[next.word] = struct{}{}
	}
	visited[i][j] = true
	for _, d := range [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		ii := i + d[0]
		jj := j + d[1]
		if ii < 0 || jj < 0 || ii >= m || jj >= n || visited[ii][jj] {
			continue
		}
		explore(board, seen, next, visited, ii, jj, m, n)
	}
	visited[i][j] = false
}
