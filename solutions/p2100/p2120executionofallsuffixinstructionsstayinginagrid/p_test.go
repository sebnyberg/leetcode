package p2120executionofallsuffixinstructionsstayinginagrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_executeInstructions(t *testing.T) {
	for _, tc := range []struct {
		n        int
		startPos []int
		s        string
		want     []int
	}{
		{3, []int{0, 1}, "RRDDLU", []int{1, 5, 4, 3, 1, 0}},
		{2, []int{1, 1}, "LURD", []int{4, 1, 0, 0}},
		{1, []int{0, 0}, "LRUD", []int{0, 0, 0, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, executeInstructions(tc.n, tc.startPos, tc.s))
		})
	}
}

func executeInstructions(n int, startPos []int, s string) []int {
	// Since there are only 500 instructions, it's easily brute-forceable (O(n^2))
	type position struct{ i, j int }
	ok := func(p position) bool {
		return p.i >= 0 && p.i < n && p.j >= 0 && p.j < n
	}
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		var moves int
		pos := position{startPos[0], startPos[1]}
		for j := i; j < len(s); j++ {
			switch s[j] {
			case 'R':
				pos.j++
			case 'D':
				pos.i++
			case 'L':
				pos.j--
			case 'U':
				pos.i--
			}
			if !ok(pos) {
				break
			}
			moves++
		}
		res[i] = moves
	}
	return res
}
