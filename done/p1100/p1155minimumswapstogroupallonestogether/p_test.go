package p1155minimumswapstogroupallonestogether

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwaps(t *testing.T) {
	for _, tc := range []struct {
		data []int
		want int
	}{
		{[]int{0, 0, 0, 1, 0}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.data), func(t *testing.T) {
			require.Equal(t, tc.want, minSwaps(tc.data))
		})
	}
}

func minSwaps(data []int) int {
	n := len(data)
	var ones int
	for _, d := range data {
		if d == 1 {
			ones++
		}
	}

	var l, r int
	var zeroes int
	for r := 0; r < ones; r++ {
		if data[r] == 0 {
			zeroes++
		}
	}
	minZeroes := zeroes
	for r < n-1 {
		if data[l] == 0 {
			zeroes--
		}
		l++
		r++
		if data[r] == 0 {
			zeroes++
		}
		if zeroes < minZeroes {
			minZeroes = zeroes
		}
	}
	return minZeroes
}
