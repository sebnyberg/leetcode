package p3147takingmaximumenergyfromthemysticdungeon

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumEnergy(t *testing.T) {
	for _, tc := range []struct {
		energy []int
		k      int
		want   int
	}{
		{[]int{5, 2, -10, -5, 1}, 3, 3},
		{[]int{-2, -3, -1}, 2, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.energy), func(t *testing.T) {
			require.Equal(t, tc.want, maximumEnergy(tc.energy, tc.k))
		})
	}
}

func maximumEnergy(energy []int, k int) int {
	res := math.MinInt32
	for i := len(energy) - 1; i >= 0; i-- {
		if i+k < len(energy) {
			energy[i] += energy[i+k]
		}
		res = max(res, energy[i])
	}
	return res
}
