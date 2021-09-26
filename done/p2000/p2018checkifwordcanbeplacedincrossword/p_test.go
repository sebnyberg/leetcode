package p2018checkifwordcanbeplacedincrossword

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_collectPlacesFromRow(t *testing.T) {
	for _, tc := range []struct {
		row   []byte
		width int
		want  [][]byte
	}{
		{[]byte("  "), 1, [][]byte{}},
		{[]byte("#   #"), 3, [][]byte{[]byte("   ")}},
		{[]byte("   #"), 3, [][]byte{[]byte("   ")}},
		{[]byte("   "), 3, [][]byte{[]byte("   ")}},
		{[]byte("#    #"), 3, [][]byte{}},
		{[]byte("#    #   "), 3, [][]byte{[]byte("   ")}},
		{[]byte("#a   # b "), 3, [][]byte{[]byte(" b ")}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.row), func(t *testing.T) {
			require.Equal(t, tc.want, findCandidates(tc.row, tc.width))
		})
	}
}

func Test_placeWordInCrossword(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{[][]byte{{'c'}, {' '}, {'#'}, {'l'}, {' '}, {'#'}, {'#'}, {' '}, {'w'}}, "c", true},
		{[][]byte{{'#', ' ', '#'}, {'#', ' ', '#'}, {'#', ' ', 'c'}}, "ca", true},
		{[][]byte{{' ', ' '}, {' ', ' '}}, "a", false},
		{[][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}, "abc", true},
		{[][]byte{{' ', '#', 'a'}, {' ', '#', 'c'}, {' ', '#', 'a'}}, "ac", false},
		{[][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', ' ', 'c'}}, "ca", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, placeWordInCrossword(tc.board, tc.word))
		})
	}
}

func placeWordInCrossword(board [][]byte, word string) bool {
	// A valid place must be surrounded by '#'
	// A valid place must be of the same length as the word.
	// A valid place must match the characters in word, or be empty.

	// Approach:
	// Scan the board for sequences of length n with '#' on both sides
	// Transpose the board
	// Scan once again for sequences of length n with '#' on both sides.
	n := len(word)
	candidates := make([][]byte, 0)
	for i := range board {
		candidates = append(candidates, findCandidates(board[i], n)...)
	}
	trans := transpose(board)
	for i := range trans {
		candidates = append(candidates, findCandidates(trans[i], n)...)
	}
	// Add each candidate reversed
	m := len(candidates)
	for i := 0; i < m; i++ {
		rev := make([]byte, 0, n)
		for j := n - 1; j >= 0; j-- {
			rev = append(rev, candidates[i][j])
		}
		candidates = append(candidates, rev)
	}

	// Iterate over the list, compare the first byte of each potential place
	// against the first byte of the word. Matches are re-sliced and put into
	// the next iteration.
	// If any matches exist at the end, we are done.
	next := make([][]byte, 0)
	for i := 0; i < n && len(candidates) > 0; i++ {
		for _, cand := range candidates {
			if cand[i] == ' ' || cand[i] == word[i] {
				next = append(next, cand)
			}
		}
		candidates, next = next, candidates
		next = next[:0]
	}

	return len(candidates) > 0
}
func transpose(board [][]byte) [][]byte {
	m, n := len(board), len(board[0])
	res := make([][]byte, n)
	for i := range res {
		res[i] = make([]byte, m)
	}
	for i := range board {
		for j := range board[i] {
			res[j][i] = board[i][j]
		}
	}
	return res
}

func findCandidates(row []byte, width int) [][]byte {
	res := make([][]byte, 0)
	for i := 0; i <= len(row)-width; i++ {
		if row[i] == '#' {
			continue
		}
		idx := bytes.IndexByte(row[i:], '#')
		if idx == -1 {
			if i+width == len(row) {
				res = append(res, row[i:])
			}
			break
		}
		if idx == width {
			res = append(res, row[i:i+width])
			i += width - 1
		} else {
			i += idx
		}
	}
	return res
}
