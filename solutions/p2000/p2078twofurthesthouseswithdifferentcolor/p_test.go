package p2078twofurthesthouseswithdifferentcolor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDistance(t *testing.T) {
	for _, tc := range []struct {
		colors []int
		want   int
	}{
		{[]int{1, 1, 1, 6, 1, 1, 1}, 3},
		{[]int{1, 8, 3, 8, 3}, 4},
		{[]int{0, 1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.colors), func(t *testing.T) {
			require.Equal(t, tc.want, maxDistance(tc.colors))
		})
	}
}

func maxDistance(colors []int) int {
	maxDist := 0
	for i := range colors {
		for j := i + 1; j < len(colors); j++ {
			d := j - i
			if colors[i] != colors[j] && d > maxDist {
				maxDist = d
			}
		}
	}
	return maxDist
}
