package p2389longestsubsequencewithlimitedsum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_answerQueries(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries []int
		want    []int
	}{
		{[]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}, []int{2, 2, 1, 1, 2, 3, 3}},
		{[]int{4, 5, 2, 1}, []int{3, 10, 21}, []int{2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, answerQueries(tc.nums, tc.queries))
		})
	}
}

func answerQueries(nums []int, queries []int) []int {
	// There are plenty of approaches to this problem.
	// The first insight is that sorting the numbers will allow us to find the
	// smallest subsequence sums in order.
	// Then you may either sort queries by size then traverse nums, or form a
	// presum list and use binary search to find the subsequence length.
	// I opted to sort queries. It requires keeping track of each queries'
	// original location.
	type queryIdx struct {
		sum int
		idx int
	}
	n := len(queries)
	qs := make([]queryIdx, n)
	for i, q := range queries {
		qs[i] = queryIdx{q, i}
	}
	sort.Ints(nums)
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].sum < qs[j].sum
	})
	res := make([]int, n)
	for i := range res {
		res[i] = len(nums)
	}
	var sum int
	var j int
	for i, x := range nums {
		for j < n && sum+x > qs[j].sum {
			res[qs[j].idx] = i
			j++
		}
		sum += x
		if j == n {
			break
		}
	}
	return res
}
