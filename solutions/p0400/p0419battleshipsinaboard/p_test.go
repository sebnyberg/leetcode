package p0419battleshipsinaboard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countBattleships(t *testing.T) {
	for _, tc := range []struct {
		board [][]byte
		want  int
	}{
		{
			[][]byte{
				{'X', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
			},
			2,
		},
		{[][]byte{{'.'}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, countBattleships(tc.board))
		})
	}
}

func countBattleships(board [][]byte) int {
	var count int
	m, n := len(board), len(board[0])
	for i := range board {
		for j, v := range board[i] {
			if v == '.' {
				continue
			}
			// Since we are scanning from top-left, we must only try right/down
			board[i][j] = '.'
			count++
			// Right
			for k := j + 1; k < n && board[i][k] == 'X'; k++ {
				board[i][k] = '.'
			}
			// Down
			for k := i + 1; k < m && board[k][j] == 'X'; k++ {
				board[k][j] = '.'
			}
		}
	}
	return count
}
