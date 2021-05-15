package p1862sumofflooredpairs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumOfFlooredPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 5, 9}, 10},
		{[]int{4, 3, 4, 3, 5}, 17},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 49},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, sumOfFlooredPairs(tc.nums))
		})
	}
}

func sumOfFlooredPairs(nums []int) int {
	sort.Ints(nums)
	var res int
	n := len(nums)
	var i int
	for i < n {
		num := nums[i]
		curPos := i
		curRes := 0
		for k := 2; curPos != n; k++ {
			pos := sort.SearchInts(nums[curPos:], k*num)
			curRes += pos * (k - 1)
			curPos += pos
		}
		count := sort.SearchInts(nums[i:], num+1)
		res += count * curRes
		i += count
	}
	return res % 1000000007
}
