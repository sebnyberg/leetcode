package p0913catandmouse

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_catMouseGame(t *testing.T) {
	for i, tc := range []struct {
		graph [][]int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[1,3],[0],[3],[0,2]]"),
			1,
		},
		{
			leetcode.ParseMatrix("[[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]"),
			0,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, catMouseGame(tc.graph))
		})
	}
}

func catMouseGame(graph [][]int) int {
	n := len(graph)
	const (
		mouse     = 0
		cat       = 1
		mouseWin  = 1 << 0
		catWin    = 1 << 1
		uncertain = 1 << 2 // uncertain outcome
		invalid   = 1 << 3 // invalid game state
	)

	// dp[i][j][k]
	// mi = mouse position
	// ci = cat position
	// who = 0 if mouse, 1 if cat
	// var dp [51][51][2]int
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, n)
	}

	// Initialize with uncertain
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = uncertain
			}
		}
	}

	// Cat is not allowed to be in position 0
	for mi := range dp {
		dp[mi][0][mouse] = invalid
		dp[mi][0][cat] = invalid
	}

	// Wherever the cat may be...
	for ci := 1; ci < n; ci++ {
		// The mouse cannot move after reaching the end (game is over)
		dp[0][ci][mouse] = invalid

		// And the mouse has won when it's the cat's turn
		dp[0][ci][cat] = mouseWin
	}

	// If the cat and mouse are in the same position, the cat has won
	for i := 1; i < n; i++ {
		dp[i][i][cat] = catWin
		dp[i][i][mouse] = catWin
	}

	for {
		var changed bool
		for mi := 0; mi < n; mi++ {
			for ci := 1; ci < n; ci++ {
				// For both the cat and mouse and an undecided state (draw),
				// explore possible next states. The cat wants to find a
				// position where the cat wins, then a draw, and finally, if all
				// else fails, it will accept a loss (mouse). The same story
				// goes for the mouse.
				if dp[mi][ci][cat] == uncertain {
					var nextState int
					for _, nei := range graph[ci] {
						nextState |= dp[mi][nei][mouse]
					}
					if nextState&catWin > 0 {
						dp[mi][ci][cat] = catWin
						changed = true
					} else if nextState&(uncertain|mouseWin) == mouseWin {
						dp[mi][ci][cat] = mouseWin
						changed = true
					}
				}
				if dp[mi][ci][mouse] == uncertain {
					var nextState int
					for _, nei := range graph[mi] {
						nextState |= dp[nei][ci][cat]
					}
					if nextState&mouseWin > 0 {
						changed = true
						dp[mi][ci][mouse] = mouseWin
					} else if nextState&(uncertain|catWin) == catWin {
						changed = true
						dp[mi][ci][mouse] = catWin
					}
				}
			}
		}
		if dp[1][2][mouse]&mouseWin > 0 {
			return 1
		} else if dp[1][2][mouse]&catWin > 0 {
			return 2
		} else if !changed {
			return 0
		}
	}
}
