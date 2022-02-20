package p0488zumagame

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMinStep(t *testing.T) {
	for _, tc := range []struct {
		board string
		hand  string
		want  int
	}{
		{"WRRWW", "BR", 1},
		{"WRRBBW", "RB", -1},
		{"WWRRBBWW", "RB", 2},
		{"WWRRBBWW", "WRBRW", 2},
		{"G", "GGGGG", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.board), func(t *testing.T) {
			require.Equal(t, tc.want, findMinStep(tc.board, tc.hand))
		})
	}
}

func findMinStep(board string, hand string) int {
	var freq [26]int8
	for i := range hand {
		freq[hand[i]-'A']++
	}
	mem := make(map[key]int, len(board))
	res := dfs(mem, board, freq)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

type key struct {
	s    string
	freq [26]int8
}

func dfs(mem map[key]int, board string, freq [26]int8) int {
	board = dedup(board)
	if board == "" {
		return 0
	}
	k := key{board, freq}
	if v, exists := mem[k]; exists {
		return v
	}
	res := math.MaxInt32

	// For each available ball in the hand
	for i, count := range freq {
		if count == 0 {
			continue
		}
		ch := byte(i + 'A')
		// If the ball is available, try adding it to every position
		n := len(board)
		for j := 0; j < n; j++ {
			// Only insert at start of sequence of same color
			if j > 0 && board[j-1] == ch {
				continue
			}
			sameColor := board[j] == ch
			betweenColors := j > 0 && board[j] == board[j-1]
			if sameColor || betweenColors {
				freq[i]--
				newBoard := board[:j] + string(byte(i+'A')) + board[j:]
				res = min(res, 1+dfs(mem, newBoard, freq))
				freq[i]++
			}
		}
	}

	mem[k] = res
	return mem[k]
}

func dedup(board string) string {
	bb := []byte(board)
	bb = append(bb, byte(0))
	var i int
	for j := 0; j < len(bb); j++ {
		if bb[i] == bb[j] {
			continue
		}
		if j-i >= 3 {
			copy(bb[i:], bb[j:])
			bb = bb[:len(bb)-(j-i)]
			j = -1
			i = 0
		} else {
			i = j
		}
	}
	return string(bb[:len(bb)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
