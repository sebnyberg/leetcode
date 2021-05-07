package p0583deleteopfortwostrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	for _, tc := range []struct {
		word1 string
		word2 string
		want  int
	}{
		{"sea", "eat", 2},
		{"leetcode", "etco", 4},
		{"dinitrophenylhydrazine", "acetylphenylhydrazine", 11},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word1), func(t *testing.T) {
			require.Equal(t, tc.want, minDistance(tc.word1, tc.word2))
		})
	}
}

func minDistance(word1 string, word2 string) int {
	h := Helper{
		mem: make(map[[2]int]int),
	}
	return h.helper([]byte(word1), 0, len(word1), []byte(word2), 0, len(word2))
}

type Helper struct {
	mem map[[2]int]int
}

func (h *Helper) helper(word1 []byte, i, n1 int, word2 []byte, j, n2 int) int {
	k := [2]int{i, j}
	if res, exists := h.mem[k]; exists {
		return res
	}
	var res int
	switch {
	// If both lens are zero, we are done
	case i == n1 && j == n2:
		res = 0
	// If one cursor is out of bounds, delete from the other
	case i == n1:
		res = 1 + h.helper(word1, i, n1, word2, j+1, n2)
	case j == n2:
		res = 1 + h.helper(word1, i+1, n1, word2, j, n2)
	// If there is a match, it is always optimal to move both cursors forward
	case word1[i] == word2[j]:
		res = h.helper(word1, i+1, n1, word2, j+1, n2)
	// If cursors mismatch, try to move each cursor forward and check which
	// one gave the best result (recurse)
	default:
		res = min(
			1+h.helper(word1, i+1, n1, word2, j, n2),
			1+h.helper(word1, i, n1, word2, j+1, n2),
		)
	}
	h.mem[k] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
