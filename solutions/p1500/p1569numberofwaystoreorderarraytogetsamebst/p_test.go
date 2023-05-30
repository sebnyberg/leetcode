package p1569numberofwaystoreorderarraytogetsamebst

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numOfWays(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{
			[]int{10, 23, 12, 18, 4, 29, 2, 8, 41, 31, 25, 21, 14, 35, 26, 5, 19, 43, 22, 37, 9, 20, 44, 28, 1, 39, 30, 38, 36, 6, 13, 16, 27, 17, 34, 7, 15, 3, 11, 24, 42, 33, 40, 32},
			182440977,
		},
		{[]int{3, 4, 5, 1, 2}, 5},
		{[]int{2, 1, 3}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numOfWays(tc.nums))
		})
	}
}

const mod = 1e9 + 7

func numOfWays(nums []int) int {
	return f(nums) - 1
}

func f(nums []int) int {
	// Split nums into "larger than" and "smaller than"
	// The number of ways to reorder the current subtree is equal to the number
	// of ways you could reorder the "larger" side times the number of ways you
	// can reorder the "smaller" side, times the number of ways you could
	// shuffle the "irrespective" slots. An "irrespective" slot is a place in
	// the array that one of the sides of the tree does not care about, or
	// rather it does not have a relationship with
	if len(nums) <= 1 {
		return 1
	}
	smaller := []int{}
	larger := []int{}
	for _, x := range nums[1:] {
		if x > nums[0] {
			larger = append(larger, x)
		}
		if x < nums[0] {
			smaller = append(smaller, x)
		}
	}
	n := len(nums) - 1
	k := len(larger)
	kn := nchoosek(n, k)
	small := f(smaller) % mod
	large := f(larger) % mod
	res := ((small * large % mod) * kn) % mod
	return res
}

func nchoosek(n, k int) int {
	if k == 0 {
		return 1
	}
	a := 1
	for x := n; x > k; x-- {
		a = (a * x) % mod
	}
	b := 1
	for x := n - k; x > 1; x-- {
		b = (b * x) % mod
	}
	inv := big.NewInt(int64(b))
	inv = inv.ModInverse(inv, big.NewInt(int64(mod)))
	modinv := inv.Int64()
	res := a * int(modinv)
	return res % mod
}

func modInverse(a, mod int) int {
	return modPow(a, mod-2, mod)
}

func modPow(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	p := modPow(a, b/2, mod) % mod
	p = p * p % mod
	if b%2 == 0 {
		return p
	}
	return (a * p) % mod
}
