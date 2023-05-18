package p1232

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_checkStraightLine(t *testing.T) {
	for i, tc := range []struct {
		coordinates [][]int
		want        bool
	}{
		{leetcode.ParseMatrix("[[0,0],[0,1],[0,-1]]"), true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, checkStraightLine(tc.coordinates))
		})
	}
}

const eps = 1e-5

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 1 {
		return true
	}
	dydx := func(i, j int) float64 {
		dx := coordinates[j][0] - coordinates[i][0]
		dy := coordinates[j][1] - coordinates[i][1]
		return float64(dy) / float64(dx)
	}
	isinf := func(x float64) bool {
		return math.IsInf(x, -1) || math.IsInf(x, 1)
	}
	want := dydx(0, 1)
	for i := 2; i < len(coordinates); i++ {
		got := dydx(0, i)
		if math.Abs(got-want) <= eps || (isinf(want) && isinf(got)) {
			continue
		}
		return false
	}
	return true
}
