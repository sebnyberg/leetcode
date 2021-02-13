package p0079wordsearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_exist(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB", false},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.board, tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, exist(tc.board, tc.word))
		})
	}
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	// Create board of blocked positions
	// Surround the original board with extra slots
	// to make it easier to check surrounding areas without
	// going out-of-bounds
	blocked := make([][]bool, m+2)
	for i := range blocked {
		blocked[i] = make([]bool, n+2)
		blocked[i][0] = true
		blocked[i][n+1] = true
	}
	for j := range blocked[0] {
		blocked[0][j] = true
		blocked[m+1][j] = true
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				blocked[i+1][j+1] = true
				if findPath(i+1, j+1, board, blocked, word[1:]) {
					return true
				}
				blocked[i+1][j+1] = false
			}
		}
	}

	return false
}

func findPath(i, j int, board [][]byte, blocked [][]bool, remainder string) bool {
	if len(remainder) == 0 {
		return true
	}

	for _, pos := range [][2]int{
		{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
	} {
		if !blocked[pos[0]][pos[1]] && board[pos[0]-1][pos[1]-1] == remainder[0] {
			blocked[pos[0]][pos[1]] = true
			if findPath(pos[0], pos[1], board, blocked, remainder[1:]) {
				return true
			}
			blocked[pos[0]][pos[1]] = false
		}
	}

	return false
}
