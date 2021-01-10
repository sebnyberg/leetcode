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
		// {
		// 	[][]byte{
		// 		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		// 		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		// 		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		// 		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		// 		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		// 		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		// 		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		// 		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		// 		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
		// 	}, true,
		// },
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, solveSudoku(tc.in))
		})
	}
}

func solveSudoku(board [][]byte) bool {
	var solver sudokuSolver
	return solver.Solve(board)
}

type cell struct {
	val      int
	options  [10]bool
	noptions int
}

type sudokuSolver struct {
	cells [9][9]cell
}

func (s *sudokuSolver) Solve(board [][]byte) bool {
	// Initialize the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.cells[i][j].options = [10]bool{false, true, true, true, true, true, true, true, true, true}
			s.cells[i][j].noptions = 9
		}
	}

	// Load preset values
	for i, row := range board {
		for j, cell := range row {
			if cell != '.' {
				s.selectValue(i, j, int(cell-'0'))
			}
		}
	}

	for s.SolveSimple() {
	}
	if !s.Backtrack(0, 0) {
		panic("failed to solve board")
	}
	// Final step - write values to the provided matrix
	for i := range board {
		for j := range board[i] {
			board[i][j] = byte('0' + s.cells[i][j].val)
		}
	}
	return true
}

func (s *sudokuSolver) Backtrack(i, j int) bool {
	if i == 9 {
		return true
	}
	if j == 9 {
		return s.Backtrack(i+1, 0)
	}
	if s.cells[i][j].val != 0 {
		return s.Backtrack(i, j+1)
	}

	for n := 1; n <= 9; n++ {
		if s.cells[i][j].options[n] {
			before := s.cells
			s.selectValue(i, j, n)
			if s.Backtrack(i, j+1) {
				return true
			}
			s.cells = before
		}
	}
	return false
}

func (s *sudokuSolver) SolveSimple() bool {
	// Look for cells where there is only one option
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.cells[i][j].val != 0 {
				continue
			}
			if s.cells[i][j].noptions == 1 {
				for k, available := range s.cells[i][j].options {
					if available {
						s.selectValue(i, j, k)
					}
				}
				return true
			}
		}
	}
	return false
}

func (s *sudokuSolver) SetOpt(i, j, val int, avail bool) {
	if s.cells[i][j].options[val] {
		s.cells[i][j].noptions--
	}
	s.cells[i][j].options[val] = false
}

// Select removes value as an option from the same row/col
// and 3x3 square by removing s.option entries.
func (s *sudokuSolver) selectValue(i, j, val int) bool {
	if !s.cells[i][j].options[val] {
		return false
	}

	for k := 0; k < 9; k++ {
		s.SetOpt(i, k, val, false)
		s.SetOpt(k, j, val, false)
	}
	// Remove from 3x3 square
	for k := (i / 3) * 3; k < (i/3)*3+3; k++ {
		for l := (j / 3) * 3; l < (j/3)*3+3; l++ {
			s.SetOpt(k, l, val, false)
		}
	}
	s.cells[i][j].val = val
	return true
}

// func (s *SudokuSolver) String() string {
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
