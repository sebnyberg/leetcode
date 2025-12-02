package p3623countnumberoftrapezoidsi

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countTrapezoids(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		want   int
	}{
		{leetcode.ParseMatrix("[[1,0],[2,0],[3,0],[2,2],[3,2]]"), 3},
		{leetcode.ParseMatrix("[[87,-39],[12,-94],[-30,-11],[-76,-11]]"), 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, countTrapezoids(tc.points))
		})
	}
}

const mod = 1e9 + 7

func countTrapezoids(points [][]int) int {
	// A valid trapezoid can be formed by any combination of two pairs of points
	// where the points share the same y-coordinate.

	// First partition by y-axis
	xs := make(map[int]int)
	for _, p := range points {
		y := p[1]
		xs[y]++
	}

	// And the number of valid points is n! / (n-r)! r! = n! / (n-2)!*2!
	// = n*(n-1)/2
	var previousPairs int
	var res int
	for _, v := range xs {
		newPairs := ((v * (v - 1)) / 2) % mod
		res = (res + (previousPairs*newPairs)%mod) % mod
		previousPairs = (previousPairs + newPairs) % mod
	}

	return res
}
