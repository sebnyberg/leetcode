package p0461hammingdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hammingDistance(t *testing.T) {
	for _, tc := range []struct {
		x    int
		y    int
		want int
	}{
		{1, 4, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.x), func(t *testing.T) {
			require.Equal(t, tc.want, hammingDistance(tc.x, tc.y))
		})
	}
}

func hammingDistance(x int, y int) int {
	if y > x {
		x, y = y, x
	}
	var dist int
	for x > 0 {
		dist += x&1 ^ y&1
		x >>= 1
		y >>= 1
	}
	return dist
}
