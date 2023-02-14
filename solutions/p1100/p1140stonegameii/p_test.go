package p1140stonegameii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_stoneGameII(t *testing.T) {
	for i, tc := range []struct {
		piles []int
		want  int
	}{
		{[]int{2, 7, 9, 4, 4}, 10},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, stoneGameII(tc.piles))
		})
	}
}

func stoneGameII(piles []int) int {
	mem := make(map[[2]int]int)
	var sum int
	for _, p := range piles {
		sum += p
	}
	res := dp(mem, piles, 0, 1, len(piles), sum)
	return res
}

// dp returns the number of points that the current player will take from this
// and all following rounds.
func dp(mem map[[2]int]int, piles []int, i, m, n, rem int) int {
	if i == n {
		return 0
	}
	k := [2]int{i, m}
	if v, exists := mem[k]; exists {
		return v
	}
	var sum int
	var res int
	for j := i; j-i+1 <= 2*m && j < n; j++ {
		sum += piles[j]
		rem -= piles[j]
		mm := max(m, j-i+1)
		other := dp(mem, piles, j+1, mm, n, rem)
		leftForMe := rem - other
		res = max(res, sum+leftForMe)
	}
	mem[k] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
