package p2275largestcombinationwithbitwiseandgreaterthanzero

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestCombination(t *testing.T) {
	for _, tc := range []struct {
		candidates []int
		want       int
	}{
		{[]int{16, 17, 71, 62, 12, 24, 14}, 4},
		{[]int{8, 8}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.candidates), func(t *testing.T) {
			require.Equal(t, tc.want, largestCombination(tc.candidates))
		})
	}
}

func largestCombination(candidates []int) int {
	var bitCount [30]int
	for _, c := range candidates {
		for i := 0; c > 0; i++ {
			if c&1 == 1 {
				bitCount[i]++
			}
			c >>= 1
		}
	}
	var res int
	for _, c := range bitCount {
		res = max(res, c)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
