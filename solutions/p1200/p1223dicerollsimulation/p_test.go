package p1223dicerollsimulation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dieSimulator(t *testing.T) {
	for i, tc := range []struct {
		n       int
		rollMax []int
		want    int
	}{
		{2, []int{1, 1, 2, 2, 2, 3}, 34},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, dieSimulator(tc.n, tc.rollMax))
		})
	}
}

const mod = 1e9 + 7

func dieSimulator(n int, rollMax []int) int {
	// Keep a count of how many sequences of rolls end with a certain number of
	// consecutive rolls.
	//
	var curr [6]map[int]int
	var next [6]map[int]int
	for i := range curr {
		curr[i] = make(map[int]int)
		next[i] = make(map[int]int)
		curr[i][1] = 1
	}

	for i := 1; i < n; i++ {
		for i := range next {
			for k := range next[i] {
				delete(next[i], k)
			}
		}
		// Two cases
		//
		// 1. Need total count of numbers for each number to calculate the
		// count of number of sequences with a single number.
		// 2. Need count of numbers ending with one less than current number for
		// all numbers <= rollMax
		var count [6]int
		var totalCount int
		for i := range curr {
			for _, cnt := range curr[i] {
				count[i] += cnt
			}
			totalCount += count[i]
		}
		for i := range curr {
			next[i][1] = (totalCount - count[i]) % mod
		}
		for i := range curr {
			for j := 1; j < rollMax[i]; j++ {
				next[i][j+1] = curr[i][j]
			}
		}

		curr, next = next, curr
	}

	var res int
	for i := range curr {
		for _, cnt := range curr[i] {
			res = (res + cnt) % mod
		}
	}
	return res
}
