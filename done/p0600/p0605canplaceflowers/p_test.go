package p0605canplaceflowers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPlaceFlowers(t *testing.T) {
	for _, tc := range []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.flowerbed), func(t *testing.T) {
			require.Equal(t, tc.want, canPlaceFlowers(tc.flowerbed, tc.n))
		})
	}
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	for i := range flowerbed {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == m-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
