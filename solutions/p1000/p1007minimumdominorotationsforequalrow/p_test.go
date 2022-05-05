package p1007minimumdominorotationsforequalrow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDominoRotations(t *testing.T) {
	for _, tc := range []struct {
		tops    []int
		bottoms []int
		want    int
	}{
		{[]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 3}, 2},
		{[]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tops), func(t *testing.T) {
			require.Equal(t, tc.want, minDominoRotations(tc.tops, tc.bottoms))
		})
	}
}

func minDominoRotations(tops []int, bottoms []int) int {
	mask := (1 << 7) - 1
	var above [7]int
	var below [7]int
	for i, v1 := range tops {
		v2 := bottoms[i]
		if v1 != v2 {
			above[v1]++
			below[v2]++
		}
		mask &= (1 << v1) | (1 << v2)
	}
	minSwaps := -1
	for i := 1; i <= 6; i++ {
		if mask&(1<<i) > 0 {
			return min(above[i], below[i])
		}
	}
	return minSwaps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
