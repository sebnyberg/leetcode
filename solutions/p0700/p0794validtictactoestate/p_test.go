package p0794validtictactoestate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validTicTacToe(t *testing.T) {
	for _, tc := range []struct {
		board []string
		want  bool
	}{
		{[]string{"XOX", "O O", "XOX"}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, validTicTacToe(tc.board))
		})
	}
}

func validTicTacToe(board []string) bool {
	// There are only 2^9 different states, so let's enumerate all of them
	states := make(map[[3][3]byte]struct{})
	var want [3][3]byte
	var state [3][3]byte
	for i := 0; i < 9; i++ {
		state[i/3][i%3] = ' '
	}
	states[state] = struct{}{}
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			want[r][c] = board[r][c]
			fill(states, state, r, c, 'X')
		}
	}
	if _, exists := states[want]; exists {
		return true
	}
	return false
}

var otherPlayer = [256]byte{'X': 'O', 'O': 'X'}

func fill(states map[[3][3]byte]struct{}, state [3][3]byte, r, c int, player byte) {
	if state[r][c] != ' ' {
		return
	}
	state[r][c] = player
	if _, exists := states[state]; exists {
		return
	}
	states[state] = struct{}{}
	// Check if we're done
	done := func() bool {
		for i := 0; i < 3; i++ {
			if state[i][0] != ' ' && state[i][0] == state[i][1] && state[i][1] == state[i][2] {
				return true
			}
			if state[0][i] != ' ' && state[0][i] == state[1][i] && state[1][i] == state[2][i] {
				return true
			}
		}
		if state[0][0] != ' ' && state[0][0] == state[1][1] && state[1][1] == state[2][2] {
			return true
		}
		if state[0][2] != ' ' && state[0][2] == state[1][1] && state[1][1] == state[2][0] {
			return true
		}
		return false
	}
	if done() {
		return
	}
	// Try all valid options
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if state[r][c] == ' ' {
				fill(states, state, r, c, otherPlayer[player])
			}
		}
	}
}
