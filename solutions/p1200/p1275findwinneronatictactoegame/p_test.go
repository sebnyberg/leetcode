package p1275findwinneronatictactoegame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_tictactoe(t *testing.T) {
	for _, tc := range []struct {
		moves [][]int
		want  string
	}{
		{[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}, "A"},
		{[][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}, "B"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.moves), func(t *testing.T) {
			require.Equal(t, tc.want, tictactoe(tc.moves))
		})
	}
}

const playerA = 1
const playerB = 2

func tictactoe(moves [][]int) string {
	var state [3][3]uint8
	for i, m := range moves {
		r, c := m[0], m[1]
		var player uint8
		if i%2 == 0 {
			player = playerA
		} else {
			player = playerB
		}
		state[r][c] = player
		if didWin(&state, player, r, c) {
			if player == playerA {
				return "A"
			} else {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func didWin(state *[3][3]uint8, player uint8, i, j int) bool {
	// Check row
	hasRow, hasCol, hasDiag, hasAntiDiag := true, true, true, true
	for k := 0; k < 3; k++ {
		if state[k][k] != player {
			hasDiag = false
		}
		if state[k][2-k] != player {
			hasAntiDiag = false
		}
		if state[i][k] != player {
			hasCol = false
		}
		if state[k][j] != player {
			hasRow = false
		}
	}
	return hasRow || hasCol || hasDiag || hasAntiDiag
}
