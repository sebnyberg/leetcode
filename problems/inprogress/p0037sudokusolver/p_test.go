package p0037sudokusolver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solveSudoku(t *testing.T) {
	for _, tc := range []struct {
		in   [][]byte
		want bool
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
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, solveSudoku(tc.in))
		})
	}
}

func solveSudoku(board [][]byte) bool {
	var row [9]int
	var col [9]int
	var block [9]int

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			v := 1 << (board[r][c] - '1')
			row[r] |= v
			col[c] |= v
			block[blockID(r, c)] |= v
		}
	}

	return solve(board, row, col, block)
}

func blockID(i, j int) int {
	return (i/3)*3 + j/3
}

func solve(board [][]byte, row, col, block [9]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				continue
			}

			opts := (row[r] | col[c] | block[blockID(r, c)]) ^ 0x1FF
			for v := 0; v < 9; v++ {
				bitval := 1 << v
				if (opts & bitval) == 0 {
					continue
				}

				row[r] |= bitval
				col[c] |= bitval
				block[blockID(r, c)] |= bitval
				board[r][c] = byte('1' + v)

				if solve(board, row, col, block) {
					return true
				}

				row[r] &^= bitval
				col[c] &^= bitval
				block[blockID(r, c)] &^= bitval
				board[r][c] = '.'
			}
			return false
		}
	}
	return true
}

// func (s *sudokuSolver) String() string {
// 	var sb strings.Builder
// 	for i, row := range s.cells {
// 		if i%3 == 0 {
// 			sb.WriteString("-------------\n")
// 		}
// 		for j, col := range row {
// 			if j%3 == 0 {
// 				sb.WriteRune('|')
// 			}
// 			if col.val == 0 {
// 				sb.WriteRune('.')
// 			} else {
// 				sb.WriteString(strconv.Itoa(col.val))
// 			}
// 		}
// 		sb.WriteRune('|')
// 		sb.WriteRune('\n')
// 	}
// 	sb.WriteString("-------------\n")
// 	sb.WriteRune('\n')
// 	return sb.String()
// }
