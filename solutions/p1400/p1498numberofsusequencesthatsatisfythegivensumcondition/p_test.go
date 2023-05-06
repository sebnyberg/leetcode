package p1498numberofsusequencesthatsatisfythegivensumcondition

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubseq(t *testing.T) {
	for i, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{3, 5, 6, 7}, 9, 4},
		{[]int{3, 3, 6, 8}, 10, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numSubseq(tc.nums, tc.target))
		})
	}
}

const mod = 1e9 + 7

func numSubseq(nums []int, target int) int {
	// The main realisation is that order does not matter.
	// We can therefore sort all values to perform easier counting.
	// Consider numbers from large to small - given that the number can be
	// matched with a certain small number to satisfy the constraint, then any
	// smaller number is also valid. For the next number, the largest small
	// number can only stay the same or increase. We do this until the two
	// indices "collide", at which point all numbers are part of each valid
	// subsequence.
	sort.Ints(nums)
	n := len(nums)
	var res int
	r := n - 1
	for l := range nums {
		for r >= l && nums[r]+nums[l] > target {
			r--
		}
		if r < l {
			break
		}
		res = (res + modpow(2, r-l, mod)) % mod
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func modpow(x, p, m int) int {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return x
	}
	half := modpow(x, p/2, m) % m
	return (half * half * modpow(x, p&1, m)) % m
}
