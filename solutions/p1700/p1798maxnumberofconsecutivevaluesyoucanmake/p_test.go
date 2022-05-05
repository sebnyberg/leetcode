package p1798maxnumberofconsecutivevaluesyoucanmake

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMaximumConsecutive(t *testing.T) {
	for _, tc := range []struct {
		coins []int
		want  int
	}{
		{[]int{1, 1, 1, 4}, 8},
		{[]int{1, 4, 10, 3, 1}, 20},
		{[]int{1, 3}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.coins), func(t *testing.T) {
			require.Equal(t, tc.want, getMaximumConsecutive(tc.coins))
		})
	}
}

func getMaximumConsecutive(coins []int) int {
	leftSum := 0
	sort.Ints(coins)
	for _, c := range coins {
		if leftSum < c-1 {
			break
		}
		leftSum += c
	}

	return leftSum + 1
}
