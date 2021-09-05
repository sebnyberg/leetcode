package p1994thenumberofgoodsubsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfGoodSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{5, 10, 1, 26, 24, 21, 24, 23, 11, 12, 27, 4, 17, 16, 2, 6, 1, 1, 6, 8, 13, 30, 24, 20, 2, 19}, 5368},
		{[]int{18, 28, 2, 17, 29, 30, 15, 9, 12}, 19},
		{[]int{1, 2, 3, 4}, 6},
		{[]int{4, 2, 3, 15}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfGoodSubsets(tc.nums))
		})
	}
}

const mod = 1e9 + 7

func numberOfGoodSubsets(nums []int) int {
	// Numbers which are prime or a product of distinct primes change the total
	// number of subsets with distinct prime products
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	// The following are not prime or distinct prime products:
	// 4, 8, 9, 12, 16, 18, 20, 24, 25, 27, 28
	// The pattern is: multiples of 4 , 9, or 25
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	dp := make([]int, 1<<len(primes)+1)
	dp[0] = 1
	for n := range count {
		if n == 1 || n%4 == 0 || n%9 == 0 || n == 25 {
			continue
		}
		var mask int
		for i, p := range primes {
			if n%p == 0 {
				mask += 1 << i
			}
		}
		for i := 0; i < (1 << len(primes)); i++ {
			if mask&i > 0 {
				continue
			}
			dp[mask|i] = (dp[mask|i] + count[n]*dp[i]) % mod
		}
	}
	var res int
	for comb, count := range dp {
		if comb == 0 {
			continue
		}
		res += count
		res %= mod
	}
	res *= powMod(2, count[1], mod)
	return res % mod
}

func powMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % mod
		}
		b >>= 1
		a = (a * a) % mod
	}
	return res
}
