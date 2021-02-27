package p0130surroundedregions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solve(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		want  [][]byte
	}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			[][]byte{{'X'}},
			[][]byte{{'X'}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			solve(tc.board)
			require.Equal(t, tc.want, tc.board)
		})
	}
}

func solve(board [][]byte) {
	n, m := len(board[0]), len(board)

	edgePosition := func(i, j int) bool {
		return i == 0 || i == m-1 || j == 0 || j == n-1
	}

	validPosition := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				continue
			}
			// Visited locations on the board are marked with 'V'
			// If the location is surrounded, it will be swapped to 'X' by the end
			// of this loop
			if board[i][j] == 'V' {
				continue
			}
			board[i][j] = 'V'

			// Find all 'O's near the current 'O'
			nearbyOs := [][2]int{{i, j}}
			var ok bool
			for i := 0; i < len(nearbyOs); i++ {
				o := nearbyOs[i]
				curI, curJ := o[0], o[1]
				// If one of the cells in this region is at the edge,
				// the whole region is OK
				ok = ok || edgePosition(curI, curJ)
				for _, pos := range [][2]int{
					{curI - 1, curJ}, // Above
					{curI + 1, curJ}, // Below
					{curI, curJ - 1}, // Left
					{curI, curJ + 1}, // Right
				} {
					if validPosition(pos[0], pos[1]) && board[pos[0]][pos[1]] == 'O' {
						board[pos[0]][pos[1]] = 'V' // Mark as visited
						nearbyOs = append(nearbyOs, pos)
					}
				}
			}
			if !ok {
				for _, o := range nearbyOs {
					board[o[0]][o[1]] = 'X'
				}
			}
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'V' {
				board[i][j] = 'O'
			}
		}
	}
}
