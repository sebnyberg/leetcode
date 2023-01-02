package p2521distinctprimefactorsofaproductofarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distinctPrimeFactors(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 4, 8, 16}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, distinctPrimeFactors(tc.nums))
		})
	}
}

func distinctPrimeFactors(nums []int) int {
	// The key property of prime numbers if that the product of the numbers
	// cannot form any new primes aside from what is already in the numbers to
	// begin with. So finding the distinct prime factors of the product is the
	// same thing as finding disinct prime factors of each number.
	factorize := func(x int) []int {
		var res []int
		for y := 2; y*y <= x; y++ {
			for x%y == 0 {
				res = append(res, y)
				x /= y
			}
		}
		if x > 1 {
			res = append(res, x)
		}
		return res
	}
	var seen [1001]bool
	var res int
	for _, x := range nums {
		if seen[x] {
			continue
		}
		ff := factorize(x)
		for _, f := range ff {
			if !seen[f] {
				res++
			}
			seen[f] = true
		}
	}
	return res
}
