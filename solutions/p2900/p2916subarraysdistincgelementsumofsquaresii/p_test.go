package p2916subarraysdistincgelementsumofsquaresii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumCounts(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1}, 15},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, sumCounts(tc.nums))
		})
	}
}

func sumCounts(nums []int) int {
	// Consider how the sum of squares of distinct counts changes for a single
	// number in the sequence of number.
	//
	// We know that there is a new sequence with a single element that is added
	// to the total result.
	//
	// For any other subarray that ends in the current number, the total delta
	// will increase only for those subarrays for which the current number did
	// not exist already.
	//
	// Keeping track of the number of subarrays that change in this manner is
	// simple, just keep track of the latest index of each number.
	//
	// The question is how the number changes the total sum. Let's say we had
	// some sequences with 2, 3, and 6 distinct numbers that now have one more
	// number. This gives us (2+1)^2 + (3+1)^2 + (6+1)^2 =
	// = (2^2 + 2*2 + 1) + (3^2 + 2*3 + 1) + (6^6 + 2*6 + 1) =
	// = (2^2 + 3^3 + 6^6) + 2*(2+3+6) + 3*1
	//
	// Hm. This is a bit messy. The first part is easy to keep track of - it's
	// already covered by the current sum of squares. The third part is also
	// easy, it's just the distance to the previous instance of the current
	// value. The middle section, however, is not straight-forward. It requires
	// us to keep track of the current number of distinct values for all
	// subarrays starting after the prior occurrence of the current value. And
	// since the current value increases those values, we need to make an update
	// as well.
	//
	// There is probably an easier way to do this, but the only way I know how
	// to do range updates like this one is to use a segment tree with lazy
	// propagation, and that is really complicated to do in the middle of a
	// competition.
	//
	// Oh well, let's do it. We'll have a segment tree containing the sum of
	// counts of distinct numbers. We'll use lazy propagation to enable
	// efficient range updates.

	const mod = 1e9 + 7

	n := 1
	m := len(nums)
	for n < m {
		n *= 2
	}
	t := make([]int, n*2)
	lazy := make([]int, n*2)

	// // push consumes a delayed lazy update and propagates the update further
	// // down the tree.
	// push := func(i, l, r int) {
	// 	if lazy[i] == 0 {
	// 		return
	// 	}
	// 	// The total sum in the current position increases by the number of
	// 	// elements times the number of delayed updates.
	// 	t[i] = (t[i] + (r-l+1)*lazy[i]) % mod
	// 	if l != r {
	// 		lazy[i*2] += lazy[i]
	// 		lazy[i*2+1] += lazy[i]
	// 	}
	// 	lazy[i] = 0
	// 	return
	// }

	// update increments the number of unique counts of the values in the range
	// [l, r] by 1 and returns the sum of unique counts.
	var update func(i, tl, tr, l, r int) int
	update = func(i, tl, tr, l, r int) int {

		// If there's a pending update, make it
		if lazy[i] != 0 {
			t[i] = (t[i] + lazy[i]*(tr-tl+1)) % mod
			if tr != tl {
				lazy[i*2] += lazy[i]
				lazy[i*2+1] += lazy[i]
			}
			lazy[i] = 0
		}

		if r < tl || l > tr {
			// no updates and no result
			return 0
		}

		if tl >= l && tr <= r {
			// update range completely contains this segment of the tree
			// perform update and return result
			t[i] = (t[i] + (tr - tl + 1)) % mod
			if tr != tl {
				lazy[i*2]++
				lazy[i*2+1]++
			}
			return t[i]
		}

		// range overlaps with one or more of the two child segments, continue
		// collecting results
		mid := tl + (tr-tl)/2

		a := update(2*i, tl, mid, l, r)
		b := update(2*i+1, mid+1, tr, l, r)

		t[i] = t[i*2] + t[i*2+1]

		return a + b
	}

	seenAt := make(map[int]int)

	// As per logic in the introduction above, for each position, we increment
	// the total sum of squares of unique counts by
	// 2*update(seenAt[nums[j]]+1, i) + (j-seenAt[nums[j]])
	var sum int
	var res int
	for j := range nums {
		i, exists := seenAt[nums[j]]
		if !exists {
			i = -1
		}
		// rangeSum contains the sum of unique counts of subsequences containing
		// the current number. We should remove (j-i) to get the sequence as it
		// were before incrementing.
		rangeSumBefore := update(1, 0, n-1, i+1, j) - (j - i)
		sum = (sum + 2*(rangeSumBefore) + (j - i)) % mod
		res = (res + sum) % mod
		seenAt[nums[j]] = j
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
