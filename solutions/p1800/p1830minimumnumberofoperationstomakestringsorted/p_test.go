package p1830minimumnumberofoperationstomakestringsorted

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeStringSorted(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"jensgfyynhtwlgnoxkkkiguizadmz", 612956436},
		{"aabaa", 2},
		{"cdbea", 63},
		{"leetcodeleetcodeleetcode", 982157772},
		{"cba", 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, makeStringSorted(tc.s))
		})
	}
}

const mod = 1000000007

func makeStringSorted(s string) int {
	// Each time rev(s[i:]) happens, s[i:] is sorted. This means that
	// the worst possible order will appear in the tail.
	//
	// The number of operations to sort the unsorted tail s[i:] is the
	// total number of permutations.
	//
	// Given a string "cccbbaabb", the total number of permutations
	// is 9! / ((repeated character count)! * ...), i.e.
	// 9! / (3! (c) * 4! (b) * 2! (a)), i.e.
	// 9! / (3!*4!*2!)
	//
	// Counting from the final number to the first, we need to calculate
	// the number of letters below the current * the number of permutations
	// for the current sequence.
	//
	var count [26]int64
	n := len(s)
	res := big.NewInt(0)
	permsTotal := big.NewInt(1)
	for i := n - 1; i >= 0; i-- {
		pos := s[i] - 'a'
		count[pos]++

		permsTotal.Mul(permsTotal, big.NewInt(int64(n-1)))
		permsTotal.Div(permsTotal, big.NewInt(count[pos]))

		newCounts := new(big.Int)
		newCounts.Mul(permsTotal, big.NewInt(sum(count[:pos])))
		newCounts.Div(newCounts, big.NewInt(int64(n-1)))

		res.Add(res, newCounts)
	}

	res.Mod(res, big.NewInt(mod))
	return int(res.Int64())
}

func sum(nums []int64) int64 {
	var count int64
	for _, n := range nums {
		count += n
	}
	return count
}
