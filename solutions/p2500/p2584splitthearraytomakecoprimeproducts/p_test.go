package p2584splitthearraytomakecoprimeproducts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findValidSplit(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{4, 7, 15, 8, 3, 5}, -1},
		{[]int{4, 7, 8, 15, 3, 5}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findValidSplit(tc.nums))
		})
	}
}

func findValidSplit(nums []int) int {
	// i is the inclusive end of the range, i.e. the range in question is [0,i],
	// which gives [0,i+1) as a Go range..
	//
	// A range is coprime iff the two ranges share no prime factors.
	// We can create a custom "product" which doesn't actually multiply the
	// numbers, but rather merges the set of unique primes. This is a
	// "destructive" operation so it should be memoized.
	//
	mem := make(map[int][]int)
	factorize := func(x int) []int {
		if v, exists := mem[x]; exists {
			return v
		}
		var res []int
		y := x
		for k := 2; k*k <= y; k++ {
			if y%k == 0 {
				res = append(res, k)
				for y%k == 0 {
					y /= k
				}
			}
		}
		if y != 1 {
			res = append(res, y)
		}
		mem[x] = res
		return res
	}
	coprime := func(a, b map[int]int) bool {
		for k, v := range a {
			if v > 0 && b[k] > 0 {
				return false
			}
		}
		return true
	}
	n := len(nums)
	right := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		for _, p := range factorize(nums[i]) {
			right[p]++
		}
	}

	left := map[int]int{}
	for i := 0; i < n-1; i++ {
		f := factorize(nums[i])
		for _, p := range f {
			right[p]--
		}
		for _, p := range f {
			left[p]++
		}
		if coprime(left, right) {
			return i
		}
	}
	return -1
}
