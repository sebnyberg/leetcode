package p0799champagnetower

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_champagneTower(t *testing.T) {
	for _, tc := range []struct {
		poured      int
		query_row   int
		query_glass int
		want        float64
	}{
		{1, 1, 1, 0},
		{2, 1, 1, 0.5},
		{25, 6, 1, 0.18750},
		{100000009, 33, 17, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.poured), func(t *testing.T) {
			require.Equal(t, tc.want, champagneTower(tc.poured, tc.query_row, tc.query_glass))
		})
	}
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	// Any overflowing glass will pour half of its excess on its left side, half
	// on its right side.
	var curr [100]float64
	curr[0] = float64(poured)
	for row := 0; row < query_row; row++ {
		var next [100]float64
		for col := 0; col < query_row; col++ {
			if curr[col] <= 1 {
				continue
			}
			half := (curr[col] - 1) / 2
			next[col] += half
			next[col+1] += half
		}
		curr = next
	}
	return math.Min(1, curr[query_glass])
}
