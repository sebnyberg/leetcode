package p0909snakesandladders

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_snakesAndLadders(t *testing.T) {
	for i, tc := range []struct {
		board [][]int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[-1,1,1,1],[-1,7,1,1],[16,1,1,1],[-1,1,9,1]]"),
			3,
		},
		{
			leetcode.ParseMatrix("[[-1,-1,-1,46,47,-1,-1,-1],[51,-1,-1,63,-1,31,21,-1],[-1,-1,26,-1,-1,38,-1,-1],[-1,-1,11,-1,14,23,56,57],[11,-1,-1,-1,49,36,-1,48],[-1,-1,-1,33,56,-1,57,21],[-1,-1,-1,-1,-1,-1,2,-1],[-1,-1,-1,8,3,-1,6,56]]"),
			4,
		},
		{leetcode.ParseMatrix("[[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]"), 4},
		{leetcode.ParseMatrix("[[-1,-1,19,10,-1],[2,-1,-1,6,-1],[-1,17,-1,19,-1],[25,-1,20,-1,-1],[-1,-1,-1,-1,15]]"), 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, snakesAndLadders(tc.board))
		})
	}
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	for i := range board {
		for j := range board {
			if board[i][j] > 0 {
				board[i][j]--
			}
		}
	}
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		board[l], board[r] = board[r], board[l]
	}
	for i := 1; i < n; i += 2 {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			board[i][l], board[i][r] = board[i][r], board[i][l]
		}
	}
	seen := make([]bool, n*n)
	curr := []int{}
	next := []int{}
	curr = append(curr, 0)
	seen[0] = true
	end := n*n - 1
	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			for d := 1; d <= 6; d++ {
				y := x + d
				if y >= n*n || seen[y] {
					continue
				}
				seen[y] = true
				i := y / n
				j := y % n
				if y == end {
					return steps
				}
				if board[i][j] == -1 {
					next = append(next, y)
					continue
				}
				y = board[i][j]
				if y == end {
					return steps
				}
				next = append(next, y)
			}
		}
		curr, next = next, curr
	}
	return -1
}
