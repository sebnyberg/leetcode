package p0529minesweeper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_updateBoard(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		click []int
		want  [][]byte
	}{
		{
			[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			[]int{3, 0},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
			[]int{1, 2},
			[][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, updateBoard(tc.board, tc.click))
		})
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	// There are four cases:
	// 1. The player clicked a mine, in which case that position can be changed
	//   and the game is over.
	r, c := click[0], click[1]
	if board[r][c] == 'M' || board[r][c] == 'X' {
		board[r][c] = 'X'
		return board
	}
	// 2. Player clicked a digit
	if board[r][c] >= '1' && board[r][c] <= '9' {
		return board
	}
	m, n := len(board), len(board[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	dirsDiag := [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	nearDiag := func(i, j int) [][2]int {
		var res [][2]int
		for _, d := range dirsDiag {
			ii := i + d[0]
			jj := j + d[1]
			if ok(ii, jj) {
				res = append(res, [2]int{ii, jj})
			}
		}
		return res
	}
	// dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// near := func(i, j int) [][2]int {
	// 	var res [][2]int
	// 	for _, d := range dirs {
	// 		ii := i + d[0]
	// 		jj := j + d[1]
	// 		if ok(ii, jj) {
	// 			res = append(res, [2]int{ii, jj})
	// 		}
	// 	}
	// 	return res
	// }
	// 3. The player clicked a place adjacent to a mine, in which case only that
	// 		cell is revealed to have the same number as number of adj. mines
	var mines int
	for _, nei := range nearDiag(r, c) {
		if board[nei[0]][nei[1]] == 'M' || board[nei[0]][nei[1]] == 'X' {
			mines++
		}
	}
	if mines > 0 {
		board[r][c] = '0' + byte(mines)
		return board
	}

	// 4. The player clicked on an empty or unrevealed empty square. Make sure
	// that all unrevealed empty and empty squares reachable from each square are
	// revealed.
	seen := make(map[[2]int]struct{})
	curr := [][2]int{{r, c}}
	next := [][2]int{}
	seen[curr[0]] = struct{}{}
	for len(curr) > 0 {
		next = next[:0]
		for _, n := range curr {
			if board[n[0]][n[1]] >= '1' && board[n[0]][n[1]] <= '9' {
				continue
			}
			var mines int
			var cands [][2]int
			for _, nei := range nearDiag(n[0], n[1]) {
				if board[nei[0]][nei[1]] == 'X' || board[nei[0]][nei[1]] == 'M' {
					mines++
					continue
				}
				if _, exists := seen[nei]; exists {
					continue
				}
				cands = append(cands, nei)
			}
			if mines > 0 {
				board[n[0]][n[1]] = '0' + byte(mines)
			} else {
				for _, c := range cands {
					seen[c] = struct{}{}
				}
				next = append(next, cands...)
				board[n[0]][n[1]] = 'B'
			}
		}

		curr, next = next, curr
	}

	return board
}
