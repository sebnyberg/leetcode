package p1231dividechocolate

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximizeSweetness(t *testing.T) {
	for _, tc := range []struct {
		sweetness []int
		k         int
		want      int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 6},
		{[]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 8, 1},
		{[]int{1, 2, 2, 1, 2, 2, 1, 2, 2}, 2, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sweetness), func(t *testing.T) {
			require.Equal(t, tc.want, maximizeSweetness(tc.sweetness, tc.k))
		})
	}
}

func maximizeSweetness(sweetness []int, k int) int {
	// The goal is to create an evenly divided cake
	// Just do binary search to find the required sweetness.
	lo, hi := 1, math.MaxInt32
	for lo < hi {
		mid := (lo + hi) / 2
		if canDivide(sweetness, mid, k+1) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

func canDivide(sweetness []int, chunkSize, k int) bool {
	var sum int
	var count int
	for i := range sweetness {
		sum += sweetness[i]
		if sum >= chunkSize {
			count++
			if count == k {
				return true
			}
			sum = 0
		}
	}
	return false
}
