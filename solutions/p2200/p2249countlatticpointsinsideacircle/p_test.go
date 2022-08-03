package p2249countlatticpointsinsideacircle

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countLatticePoints(t *testing.T) {
	for _, tc := range []struct {
		circles [][]int
		want    int
	}{
		{
			leetcode.ParseMatrix("[[2,2,1]]"),
			5,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.circles), func(t *testing.T) {
			require.Equal(t, tc.want, countLatticePoints(tc.circles))
		})
	}
}

func countLatticePoints(circles [][]int) int {
	seen := make(map[[2]int]struct{})
	for _, c := range circles {
		x, y, r := c[0], c[1], c[2]
		for xx := x - r; xx <= x+r; xx++ {
			dx := abs(xx - x)
			for yy := y - r; yy <= y+r; yy++ {
				dy := abs(yy - y)
				if dx*dx+dy*dy <= r*r {
					seen[[2]int{xx, yy}] = struct{}{}
				}
			}
		}
	}
	return len(seen)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
