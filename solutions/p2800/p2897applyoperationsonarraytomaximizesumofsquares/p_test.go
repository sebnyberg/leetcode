package p2897applyoperationsonarraytomaximizesumofsquares

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSum(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 6, 5, 8}, 2, 261},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxSum(tc.nums, tc.k))
		})
	}
}

const mod = 1e9 + 7

func maxSum(nums []int, k int) int {
	var bitCount [32]int
	for _, x := range nums {
		var i int
		for x > 0 {
			bitCount[i] += x & 1
			x >>= 1
			i++
		}
	}

	var res int
	for ; k > 0; k-- {
		var y int
		for i := 0; i < 32; i++ {
			if bitCount[i] > 0 {
				y += (1 << i)
				bitCount[i]--
			}
		}
		res = (res + y*y) % mod
	}
	return res
}
