package p2563countthenumberoffairpairs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countFairPairs(t *testing.T) {
	for i, tc := range []struct {
		nums  []int
		lower int
		upper int
		want  int64
	}{
		{[]int{0, 1, 7, 4, 4, 5}, 3, 6, 6},
		{[]int{1, 7, 9, 2, 5}, 11, 11, 1},
		{[]int{0, 0, 0, 0, 0, 0}, 0, 0, 15},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countFairPairs(tc.nums, tc.lower, tc.upper))
		})
	}
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	// For each number, find range of valid counterparts.
	//
	// Reduce total by half and ensure not to double count the number itself.
	//
	sort.Ints(nums)
	n := len(nums)
	var res int64
	for _, x := range nums {
		wantLo := lower - x
		i := sort.SearchInts(nums, wantLo)
		if i >= n {
			continue
		}
		wantHi := upper - x
		j := sort.SearchInts(nums, wantHi+1)
		j -= 1
		if j-i < 0 {
			continue
		}
		if x >= wantLo && x <= wantHi {
			res-- // avoid matching a number with itselt
		}
		res += int64(j - i + 1)
	}
	res /= 2 // symmetric pairs
	return res
}
