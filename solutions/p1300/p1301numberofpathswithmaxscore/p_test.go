package p1301numberofpathswithmaxscore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pathsWithMaxScore(t *testing.T) {
	for i, tc := range []struct {
		board []string
		want  []int
	}{
		{[]string{"E11", "XXX", "11S"}, []int{0, 0}},
		{[]string{"E23", "2X2", "12S"}, []int{7, 1}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, pathsWithMaxScore(tc.board))
		})
	}
}

const mod = 1e9 + 7

func pathsWithMaxScore(board []string) []int {
	n := len(board[0])
	m := len(board)
	below := make([][2]int, n+1)
	curr := make([][2]int, n+1)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				curr[j] = [2]int{0, 1}
				continue
			}
			if board[i][j] == 'X' {
				curr[j] = [2]int{0, 0}
				continue
			}
			var val int
			var nways int
			for _, x := range [][2]int{
				below[j+1], below[j], curr[j+1],
			} {
				if x[0] < val {
					continue
				}
				if x[0] > val {
					val = x[0]
					nways = x[1]
					continue
				}
				nways = (nways + x[1]) % mod
			}
			if nways == 0 {
				curr[j] = [2]int{0, 0}
				continue
			}
			curr[j] = [2]int{val, nways}
			if board[i][j] != 'E' {
				curr[j][0] += int(board[i][j] - '0')
			}
		}
		curr, below = below, curr
	}
	return []int{below[0][0], below[0][1]}
}
