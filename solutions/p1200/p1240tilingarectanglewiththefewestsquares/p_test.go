package p1240tilingarectanglewiththefewestsquares

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_tilingRectangle(t *testing.T) {
	for i, tc := range []struct {
		n    int
		m    int
		want int
	}{
		{2, 3, 3},
		{5, 8, 5},
		{11, 13, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, tilingRectangle(tc.n, tc.m))
		})
	}
}

func tilingRectangle(n int, m int) int {
	// Let's try exhaustive search

	// Current state is a list of heights for each x-value
	var state [13]uint8

	// Once all heights = m, we are done
	var done [13]uint8
	for i := 0; i < n; i++ {
		done[i] = uint8(m)
	}

	seen := make(map[[13]uint8]struct{})
	res := math.MaxInt32
	dfs(seen, state, done, m, n, 0, &res)
	return res
}

func dfs(seen map[[13]uint8]struct{}, state, done [13]uint8, m, n, cur int, res *int) {
	if cur >= *res {
		return
	}
	if state == done {
		*res = min(*res, cur)
	}
	if _, exists := seen[state]; exists {
		return
	}
	seen[state] = struct{}{}

	// find bottom-most, left-most location
	minY := math.MaxInt32
	var startI int
	for i := 0; i < n; i++ {
		if int(state[i]) == m || int(state[i]) >= minY {
			continue
		}
		minY = int(state[i])
		startI = i
	}

	// check largest valid square to place
	j := startI
	for ; j+1 < n && int(state[j+1]) == minY; j++ {
	}

	for k := j; k >= startI; k-- {
		// place square
		width := k - startI + 1
		if int(state[startI])+width > m {
			continue
		}
		cpy := state
		for m := k; m >= startI; m-- {
			cpy[m] += uint8(width)
		}
		dfs(seen, cpy, done, m, n, cur+1, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
