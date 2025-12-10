package p3577countthenumberofcomputerunlockingpermutations

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPermutations(t *testing.T) {
	for _, tc := range []struct {
		complexity []int
		want       int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{3, 3, 3, 4, 4, 4}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.complexity), func(t *testing.T) {
			require.Equal(t, tc.want, countPermutations(tc.complexity))
		})
	}
}

const mod = 1e9 + 7

func countPermutations(complexity []int) int {
	// Basically, the result is the number of permutations of values after the first
	minVal := math.MaxInt32
	for _, x := range complexity[1:] {
		minVal = min(minVal, x)
	}
	if minVal <= complexity[0] {
		return 0
	}

	n := len(complexity)
	res := 1
	for k := n - 1; k >= 2; k-- {
		res = (res * k) % mod
	}
	return res
}
