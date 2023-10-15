package p2902countofsubmultisetswithboundedsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubMultisets(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		l    int
		r    int
		want int
	}{
		{[]int{1, 2, 2, 3}, 6, 6, 1},
		{[]int{1, 2, 1, 3, 5, 2}, 3, 5, 9},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countSubMultisets(tc.nums, tc.l, tc.r))
		})
	}
}

const mod = 1e9 + 7

func countSubMultisets(nums []int, l int, r int) int {
	// The sub-multiset is simply a subset of values from nums. Jeez why did
	// Leetcode use such a shitty "formal" definition.
	//
	// The problem does not state that it is looking for unique subsets either,
	// but the examples appear to applapply such a constraint.
	//
	// So we want to find the unique combination of counts of elements from nums
	// such that l <= sum(set) <= r.
	//
	// Because 0 <= l <= r <= 2*10^4, we know that the total state-space for a
	// given iteration will never exceed 2*10^4. Any prior sums above r are
	// uninteresting.
	//
	// Because the results is about unique elements, the original ordering does
	// not matter. Also, we can take advantage of the fact that adding a unique
	// count of numbers is guaranteed to result in a unique multiset.
	//
	// The solution below is O(n^2), which appeared to be OK for Go.
	//

	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}

	curr := make([]int, r+1)
	curr[0] = 1
	next := make([]int, r+1)
	for x, cnt := range count {
		copy(next, curr)
		for k := 1; k <= cnt; k++ {
			y := k * x
			for i := 0; i+y <= r; i++ {
				next[i+y] = (next[i+y] + curr[i]) % mod
			}
		}
		curr, next = next, curr
	}
	var res int
	for i := l; i <= r; i++ {
		res = (res + curr[i]) % mod
	}

	return res
}
