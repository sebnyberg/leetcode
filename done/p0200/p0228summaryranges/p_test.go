package p0228summaryranges

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_summaryRanges(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []string
	}{
		{[]int{1, 2, 4, 5, 7}, []string{"1->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
		{[]int{-1}, []string{"-1"}},
		{[]int{0}, []string{"0"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, summaryRanges(tc.nums))
		})
	}
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	sort.Ints(nums)
	res := make([]string, 0)
	start := 0
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i]-nums[i-1] > 1 {
			// flush previous range
			if i-1 != start {
				res = append(res, fmt.Sprintf("%v->%v", nums[start], nums[i-1]))
			} else {
				res = append(res, strconv.Itoa(nums[i-1]))
			}
			start = i
		}
	}
	// Flush last range
	if start < len(nums)-1 {
		res = append(res, fmt.Sprintf("%v->%v", nums[start], nums[len(nums)-1]))
	} else {
		res = append(res, strconv.Itoa(nums[start]))
	}
	return res
}
