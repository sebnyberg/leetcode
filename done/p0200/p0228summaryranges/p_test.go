package p0228summaryranges

import (
	"fmt"
	"math"
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
		{[]int{}, nil},
		{[]int{-1}, []string{"-1"}},
		{[]int{0}, []string{"0"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, summaryRanges(tc.nums))
		})
	}
}

func summaryRanges(nums []int) []string {
	nums = append(nums, math.MaxInt64)
	stack := []int{}
	var res []string
	for _, num := range nums {
		if len(stack) > 0 && stack[len(stack)-1] != num-1 {
			if len(stack) == 1 {
				res = append(res, fmt.Sprint(stack[0]))
			} else {
				res = append(res, fmt.Sprintf("%v->%v", stack[0], stack[len(stack)-1]))
			}
			stack = stack[:0]
		}
		stack = append(stack, num)
	}
	return res
}
