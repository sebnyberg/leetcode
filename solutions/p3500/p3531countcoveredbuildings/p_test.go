package p3561countcoveredbuildings

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countCoveredBuildings(t *testing.T) {
	for _, tc := range []struct {
		n         int
		buildings [][]int
		want      int
	}{
		{3, leetcode.ParseMatrix("[[1,2],[2,2],[3,2],[2,1],[2,3]]"), 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countCoveredBuildings(tc.n, tc.buildings))
		})
	}
}

func countCoveredBuildings(n int, buildings [][]int) int {
	// First, let's count points per x- coordinate
	xRight := make(map[int]int)
	yAbove := make(map[int]int)
	for _, b := range buildings {
		x := b[0]
		y := b[1]
		xRight[x]++
		yAbove[y]++
	}

	// Then let's sort by x-axis then y-axis
	sort.Slice(buildings, func(i, j int) bool {
		a := buildings[i]
		b := buildings[j]
		x0 := a[0]
		y0 := a[1]
		x1 := b[0]
		y1 := b[1]
		if x0 == x1 {
			return y0 < y1
		}
		return x0 < x1
	})

	xLeft := make(map[int]int)
	yBelow := make(map[int]int)
	var res int
	for _, p := range buildings {
		x := p[0]
		y := p[1]

		// Remove this point from the counts
		xRight[x]--
		yAbove[y]--

		// Check if this point is surrounded
		if xLeft[x] > 0 && xRight[x] > 0 && yBelow[y] > 0 && yAbove[y] > 0 {
			res++
		}

		// Add this point to the xLeft / yBelow
		xLeft[x]++
		yBelow[y]++

	}
	return res
}
