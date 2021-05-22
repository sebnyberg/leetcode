package p0348designtictactoe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTicTacToe(t *testing.T) {
	b := Constructor(3)
	for _, tc := range []struct {
		row, col, player, want int
	}{
		// {0, 0, 1, 0},
		// {0, 2, 2, 0},
		// {2, 2, 1, 0},
		// {1, 1, 2, 0},
		// {2, 0, 1, 0},
		// {1, 0, 2, 0},
		// {2, 1, 1, 1},
		{1, 1, 1, 0},
		{0, 1, 2, 0},
		{0, 0, 1, 0},
		{2, 2, 2, 0},
		{0, 2, 1, 0},
		{2, 0, 2, 0},
		{1, 0, 1, 0},
		{2, 1, 2, 2},
	} {
		t.Run(fmt.Sprintf("%v,%v,%v,%v", tc.row, tc.col, tc.player, tc.want), func(t *testing.T) {
			res := b.Move(tc.row, tc.col, tc.player)
			require.Equal(t, tc.want, res)
		})
	}
	// finalRes := b.Move(2, 1, 1)
	// require.Equal(t, 1, finalRes)
}

type TicTacToe struct {
	rows  [][2]int
	cols  [][2]int
	diags [2][2]int
	n     int
	win   int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	t := TicTacToe{
		rows: make([][2]int, n),
		cols: make([][2]int, n),
		n:    n,
		win:  (1 << n) - 1,
	}
	return t
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	// Diagonal
	rowMark := 1 << row
	if row == col {
		this.diags[0][player-1] |= rowMark
		if this.diags[0][player-1] == this.win {
			return player
		}
	}
	if row == this.n-col-1 {
		this.diags[1][player-1] |= rowMark
		if this.diags[1][player-1] == this.win {
			return player
		}
	}
	// Row
	this.rows[row][player-1] |= 1 << col
	if this.rows[row][player-1] == this.win {
		return player
	}

	// Col
	this.cols[col][player-1] |= rowMark
	if this.cols[col][player-1] == this.win {
		return player
	}
	return 0
}
