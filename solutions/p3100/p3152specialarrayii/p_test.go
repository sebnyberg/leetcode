package p3152specialarrayii

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_isArraySpecial(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		want    []bool
	}{
		{[]int{4, 3, 1, 6}, leetcode.ParseMatrix("[[0,2],[2,3]]"), []bool{false, true}},
		{[]int{2, 1}, leetcode.ParseMatrix("[[0,1]]"), []bool{true}},
		{[]int{3, 4, 1, 2, 6}, leetcode.ParseMatrix("[[0,4]]"), []bool{false}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, isArraySpecial(tc.nums, tc.queries))
		})
	}
}

func isArraySpecial(nums []int, queries [][]int) []bool {
	// There are many options here, the problem is finding an easy solution.
	// One that I can think of is to count evens/odds for each position, forming
	// two prefix sums for even/odd positions. That way, it's possible to check
	// any range.
	n := len(nums)
	oddsAtEvens := make([]int, n+1)
	oddsAtOdds := make([]int, n+1)
	for i := range nums {
		if i%2 == 0 {
			oddsAtEvens[i+1] = oddsAtEvens[i] + nums[i]&1
			oddsAtOdds[i+1] = oddsAtOdds[i]
		} else {
			oddsAtOdds[i+1] = oddsAtOdds[i] + nums[i]&1
			oddsAtEvens[i+1] = oddsAtEvens[i]
		}
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		d := q[1] - q[0] + 1
		oddsAtOdd := oddsAtOdds[q[1]+1] - oddsAtOdds[q[0]]
		oddsAtEven := oddsAtEvens[q[1]+1] - oddsAtEvens[q[0]]
		outer := d/2 + d&1
		inner := d / 2
		if q[0]%2 == 0 {
			// outer are evens, inner are odds
			// there are two cases,
			// either all outers are even and all inner are odd
			res[i] = oddsAtEven == 0 && oddsAtOdd == inner ||
				oddsAtOdd == 0 && oddsAtEven == outer
		} else {
			// outer are odds, inner are evens
			res[i] = oddsAtOdd == 0 && oddsAtEven == inner ||
				oddsAtEven == 0 && oddsAtOdd == outer
		}
	}
	return res
}
