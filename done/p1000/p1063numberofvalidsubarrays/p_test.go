package p1063numberofvalidsubarrays

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 2, 5, 3}, 11},
		{[]int{3, 2, 1}, 3},
		{[]int{2, 2, 2}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, validSubarrays(tc.nums))
		})
	}
}

func validSubarrays(nums []int) int {
	stack := list.New()
	res := 0
	for _, n := range nums {
		for stack.Len() > 0 && n < stack.Back().Value.(int) {
			stack.Remove(stack.Back())
		}
		stack.PushBack(n)
		res += stack.Len()
	}
	return res
}
