package p0575distributecandies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distributeCandies(t *testing.T) {
	for _, tc := range []struct {
		candyType []int
		want      int
	}{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
		{[]int{6, 6, 6, 6}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.candyType), func(t *testing.T) {
			require.Equal(t, tc.want, distributeCandies(tc.candyType))
		})
	}
}
func distributeCandies(candyType []int) int {
	var ntypes int
	candyTypes := make(map[int]struct{}, len(candyType)/2)
	for _, t := range candyType {
		if _, exists := candyTypes[t]; !exists {
			candyTypes[t] = struct{}{}
			ntypes++
		}
	}
	maxCandies := len(candyType) / 2
	if ntypes < maxCandies {
		return ntypes
	}
	return maxCandies
}
