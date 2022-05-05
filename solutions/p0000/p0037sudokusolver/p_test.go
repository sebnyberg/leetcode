package p0037sudokusolver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solveSudoku(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		want  bool
	}{
		{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}, true,
		},
		{
			[][]byte{
				{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
				{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
				{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
				{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
				{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
				{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
			}, true,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			solveSudoku(tc.board)
			require.True(t, isValidSudoku(tc.board))
		})
	}
}

func solveSudoku(board [][]byte) {
	var rows [9]int
	var cols [9]int
	var blocks [9]int

	// Initialize board
	for i := range board {
		for j, val := range board[i] {
			b := 1 << (val - '0')
			rows[i] |= b
			cols[j] |= b
			blocks[blockIdx(i, j)] |= b
		}
	}

	findSolution(board, 0, 0, &rows, &cols, &blocks)
}

func blockIdx(i, j int) int {
	return 3*(i/3) + j/3
}

// findSolution works with the board to find the solution (a solution is
// guaranteed by the problem). If the solution was found, it returns true,
// which is used to stop early.
func findSolution(board [][]byte, i, j int, rows, cols, blocks *[9]int) bool {
	if j == 9 {
		j = 0
		i += 1
	}
	if i == 9 {
		return true
	}
	if board[i][j] != '.' {
		return findSolution(board, i, j+1, rows, cols, blocks)
	}
	// Place each possible number
	k := blockIdx(i, j)
	for a := 1; a <= 9; a++ {
		bit := 1 << a
		if rows[i]&bit > 0 || cols[j]&bit > 0 || blocks[k]&bit > 0 {
			continue
		}
		// Try this number
		board[i][j] = byte('0' + a)
		rows[i] |= bit
		cols[j] |= bit
		blocks[k] |= bit
		if findSolution(board, i, j+1, rows, cols, blocks) {
			return true
		}
		// Undo selected number
		board[i][j] = '.'
		rows[i] &^= bit
		cols[j] &^= bit
		blocks[k] &^= bit
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	var rows [9]uint16
	var cols [9]uint16
	var boxes [9]uint16
	for row := range board {
		for col := range board[row] {
			val := int(board[row][col] - '0')
			boxIdx := 3*(row/3) + col/3
			var bit uint16 = 1 << val
			if rows[row]&bit > 0 ||
				cols[col]&bit > 0 ||
				boxes[boxIdx]&bit > 0 {
				return false
			}
			rows[row] |= bit
			cols[col] |= bit
			boxes[boxIdx] |= bit
		}
	}
	return true
}
