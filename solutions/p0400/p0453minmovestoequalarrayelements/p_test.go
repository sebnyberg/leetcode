package p0453minmovestoequalarrayelements

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMoves(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minMoves(tc.nums))
		})
	}
}

func minMoves(nums []int) int {
	minVal := math.MaxInt32
	res := 0
	for _, n := range nums {
		res += n
		if n < minVal {
			minVal = n
		}
	}
	return res - len(nums)*minVal
}
