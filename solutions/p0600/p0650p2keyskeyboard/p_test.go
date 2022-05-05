package p0650p2keyskeyboard

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSteps(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{9, 6},
		{7, 7},
		{3, 3},
		{1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minSteps(tc.n))
		})
	}
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	mem := make(map[[2]int16]int16)
	res := 1 + dfs(mem, 1, int16(n), 1)
	return int(res)
}

func dfs(mem map[[2]int16]int16, cur, n, copy int16) int16 {
	k := [2]int16{cur, copy}
	if v, exists := mem[k]; exists {
		return v
	}
	if cur == n {
		return 0
	}
	if cur > n {
		return math.MaxInt16 - 1001
	}
	keepGoing := 1 + dfs(mem, cur+copy, n, copy)
	copyPaste := 2 + dfs(mem, cur+cur, n, cur)
	res := min(keepGoing, copyPaste)
	mem[k] = res
	return res
}

func min(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}
