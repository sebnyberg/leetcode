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
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			solve(tc.board)
			require.Equal(t, tc.want, tc.board)
		})
	}
}

func solve(board [][]byte) {
	// A mark on the board is not surrounded if it:
	m, n := len(board), len(board[0])
	surrounded := make([][]bool, m+2)
	for i := range surrounded {
		surrounded[i] = make([]bool, n+2)
	}

	// All 'X'es on the board are surrounded
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				surrounded[i+1][j+1] = true
			}
		}
	}

	//
}
