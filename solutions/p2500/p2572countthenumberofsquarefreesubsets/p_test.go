package p2572countthenumberofsquarefreesubsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_squareFreeSubsets(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{22, 2, 5, 26, 28, 8, 4, 11, 12, 17, 11, 3, 19, 29, 19, 7, 24, 12, 22, 5, 8, 22}, 1727},
		{[]int{11, 2, 19, 7, 9, 27}, 15},
		{[]int{26, 6, 4, 27, 6, 18}, 3},
		{[]int{3, 4, 4, 5}, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, squareFreeSubsets(tc.nums))
		})
	}
}

const mod = 1e9 + 7

func squareFreeSubsets(nums []int) int {
	// This is a tricky combinatorics problem.
	//
	// * Any non-1s cannot be combined with itself, or else that set would
	// contain a square.
	// * Any number that is a square is invalid
	// * Any number that is a multiple of a square is invalid
	// * 1s must be handled separately from other numbers
	// * There are not that many valid numbers in the range [2,30], so the total
	// number of possible sets is not that large, perhaps around 2^12
	// * Because each (non-1) number can only exist at most once, a valid set
	// can be represented by a 32-bit integer.
	//
	// This gives us the solution: factorize each number into a bitset. If the
	// factorization is valid, it can be added on its own (1) and to any prior
	// disjoint factorized set.
	//
	// Finally, there are 2^count[1] - 1 sets which contain 1s, these can both
	// be added on their own or combined with prior sets.
	//
	// Be careful with mod.

	// factorize x to a bitset
	fac := func(x int) int {
		var res int
		for k := 2; k <= x; k++ {
			if x%(k*k) == 0 {
				return -1
			}
			if x%k == 0 {
				res |= (1 << k)
			}
		}
		return res
	}

	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}

	comb := make(map[int]int)
	for x, c1 := range count {
		if x == 1 {
			// 1 needs special handling
			continue
		}
		a := fac(x)
		if a < 0 {
			continue
		}
		comb[a] = (comb[a] + c1) % mod
		for b, c2 := range comb {
			if a&b == 0 {
				comb[a|b] = (comb[a|b] + c1*c2) % mod
			}
		}
	}

	var res int
	for _, c := range comb {
		res = (res + c) % mod
	}

	if count[1] > 0 {
		x := modpow(2, count[1], mod)
		x = (x + mod - 1) % mod
		res = (res + x*res + x) % mod
	}

	return res % mod
}

func modpow(a, b, m int) int {
	if b == 0 {
		return 1
	}
	l := modpow(a, b/2, m) % m
	if b&1 == 1 {
		return l * l * a % m
	}
	return l * l % mod
}
