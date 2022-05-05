package p1962removestonestominimizethetotal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minStoneSum(t *testing.T) {
	for _, tc := range []struct {
		piles []int
		k     int
		want  int
	}{
		{[]int{10000}, 10000, 1},
		{[]int{5, 4, 9}, 2, 12},
		{[]int{4, 3, 6, 7}, 3, 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.piles), func(t *testing.T) {
			require.Equal(t, tc.want, minStoneSum(tc.piles, tc.k))
		})
	}
}

func minStoneSum(piles []int, k int) int {
	var res int
	var factorFreq [5001]int
	for _, p := range piles {
		res += p
		for p > 1 {
			factorFreq[p/2]++
			p -= p / 2
		}
	}
	for divisor := 5000; divisor >= 0; divisor-- {
		count := factorFreq[divisor]
		if count == 0 {
			continue
		}
		d := min(k, count)
		res -= d * divisor
		k -= d
		if k == 0 {
			break
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
