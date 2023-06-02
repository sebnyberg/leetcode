package p1659maximizegridhappiness

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMaxGridHappiness(t *testing.T) {
	for i, tc := range []struct {
		m          int
		n          int
		introverts int
		extroverts int
		want       int
	}{
		{3, 1, 2, 1, 260},
		{2, 2, 4, 0, 240},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getMaxGridHappiness(tc.n, tc.m, tc.introverts, tc.extroverts))
		})
	}
}

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	var change [3][3]int
	change[introvert][introvert] = -60
	change[introvert][extrovert] = -30 + 20
	change[extrovert][introvert] = -30 + 20
	change[extrovert][extrovert] = 40
	mem := make(map[state]int)
	return dp(mem, &change, 0, 0, 0, introvertsCount, extrovertsCount, m, n)
}

const (
	nobody    = 0
	introvert = 1
	extrovert = 2
)

func dp(mem map[state]int, change *[3][3]int, prevRow, row, col, ints, exts, m, n int) int {
	if col == n {
		col = 0
		row++
	}
	if row == m {
		return 0
	}
	prevRow &= (1 << (n*2 + 2)) - 1 // clear any slot older than 5 steps back
	key := state{prevRow, row, col, ints, exts}
	if v, exists := mem[key]; exists {
		return v
	}
	above := (prevRow & (3 << (n * 2))) >> (n * 2)
	var left int
	if col > 0 {
		left = prevRow & (3 << 2) >> 2
	}
	// Try doing nothing
	res := dp(mem, change, prevRow<<2, row, col+1, ints, exts, m, n)

	// Or place an introvert
	if ints > 0 {
		val := 120 + change[introvert][above] + change[introvert][left]
		res = max(res, val+dp(mem, change, (prevRow+introvert)<<2, row, col+1, ints-1, exts, m, n))
	}

	if exts > 0 {
		val := 40 + change[extrovert][above] + change[extrovert][left]
		res = max(res, val+dp(mem, change, (prevRow+extrovert)<<2, row, col+1, ints, exts-1, m, n))
	}
	mem[key] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type state struct {
	prevRow int
	row     int
	col     int
	ints    int
	exts    int
}
