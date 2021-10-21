package p1931paintingagridwiththreedifferentcolors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_colorTheGrid(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		want int
	}{
		{1, 1, 3},
		{1, 2, 6},
		{5, 5, 580986},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, colorTheGrid(tc.m, tc.n))
		})
	}
}

type color int

const (
	red color = iota
	blue
	green
)

var rowColoring int

func colorTheGrid(m int, n int) int {
	var dp [1024][1024]int
	helper(&dp, 0, 0, 0, 0, m, n)
}

func helper(dp *[1024][1024]int, cur, above, i, j, m, n int) int {
	// Result == number of ways in which the color to the left and above is not
	// the current color.
	for color := 0; color < 3; color++ {
		if i == 0 {
			// Only care about the color above

		}
	}
}
