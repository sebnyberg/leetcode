package p1958checkifmoveislegal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkMove(t *testing.T) {
	for _, tc := range []struct {
		board        [][]byte
		rMove, cMove int
		color        byte
		want         bool
	}{
		{[][]byte{
			{'B', 'W', '.', 'B', 'W', 'W', 'B', '.'},
			{'B', '.', '.', 'B', 'W', 'W', '.', '.'},
			{'W', 'W', '.', 'B', 'B', '.', 'B', 'W'},
			{'B', 'W', 'B', '.', 'B', '.', 'B', 'B'},
			{'B', 'W', 'W', 'B', '.', 'W', 'B', 'B'},
			{'W', 'W', '.', 'B', 'W', 'B', '.', '.'},
			{'W', '.', 'B', 'W', 'W', 'B', '.', 'B'},
			{'W', '.', 'B', 'B', '.', 'B', '.', '.'}},
			2, 5, 'B', true,
		},
		{[][]byte{{'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}},
			4, 3, 'B', true,
		},
		{[][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'B', '.', '.', 'W', '.', '.', '.'}, {'.', '.', 'W', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'W', 'B', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', 'B', 'W', '.', '.'}, {'.', '.', '.', '.', '.', '.', 'W', '.'}, {'.', '.', '.', '.', '.', '.', '.', 'B'}},
			4, 4, 'W', false,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, checkMove(tc.board, tc.rMove, tc.cMove, tc.color))
		})
	}
}

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	// A good line consists of at least three cells, vertical, diagonal or
	// horizontal, where the two end-point cells share the same color.
	var isBlack bool
	if color == 'B' {
		isBlack = true
	}

	// Starting in [rMove,cMove], try all directions and check if there is a
	// good line starting with that position.
	for r := -1; r <= 1; r++ {
		for l := -1; l <= 1; l++ {
			if r == 0 && l == 0 {
				continue
			}
			if hasLine(board, len(board), len(board[0]), rMove, cMove, r, l, isBlack) {
				return true
			}
		}
	}
	return false
}

func hasLine(board [][]byte, m, n, i, j, di, dj int, isBlack bool) bool {
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && board[i][j] != '.'
	}

	// First cell must be the opposite of the marked color
	i += di
	j += dj
	want := []byte{'B', 'W'}
	if isBlack {
		want = []byte{'W', 'B'}
	}
	if !ok(i, j) || board[i][j] == want[1] {
		return false
	}
	i += di
	j += dj
	// Keep parsing while ok and opposite color
	for ok(i, j) && board[i][j] == want[0] {
		i += di
		j += dj
	}
	return ok(i, j) && board[i][j] == want[1]
}
