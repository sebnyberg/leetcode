package p2038removecoloredpiecesifbothneighboursarethesamecolor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_winnerOfGame(t *testing.T) {
	for _, tc := range []struct {
		colors string
		want   bool
	}{
		{"AAABABB", true},
		{"AA", false},
		{"ABBBBBBBAAA", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.colors), func(t *testing.T) {
			require.Equal(t, tc.want, winnerOfGame(tc.colors))
		})
	}
}

func winnerOfGame(colors string) bool {
	// Any sequence with three or more of a certain letter results in
	// len(seq) - 2 moves for that player. Then it's just a matter of counting
	// how many moves each player has to find the winner.
	count := 1
	colors += "C" // sentinel
	var aliceMoves, bobMoves int
	for i := 1; i < len(colors); i++ {
		if colors[i-1] == colors[i] {
			count++
		} else {
			if count >= 3 {
				if colors[i-1] == 'A' {
					aliceMoves += count - 2
				} else {
					bobMoves += count - 2
				}
			}
			count = 1
		}
	}
	return aliceMoves > bobMoves
}
