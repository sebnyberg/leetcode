package p2186minnumberofstepstomaketwostringsanagram2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSteps(t *testing.T) {
	for _, tc := range []struct {
		s    string
		t    string
		want int
	}{
		{"leetcode", "coats", 7},
		{"night", "thing", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minSteps(tc.s, tc.t))
		})
	}
}

func minSteps(s string, t string) int {
	var counts [2][26]int
	for _, ch := range s {
		counts[0][ch-'a']++
	}
	for _, ch := range t {
		counts[1][ch-'a']++
	}
	var moves int
	for i := range counts[0] {
		moves += abs(counts[0][i] - counts[1][i])
	}
	return moves
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
