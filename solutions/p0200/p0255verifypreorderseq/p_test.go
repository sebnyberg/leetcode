package p0255verifypreorderseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_verifyPreorder(t *testing.T) {
	for _, tc := range []struct {
		preorder []int
		want     bool
	}{
		{[]int{5, 2, 1, 3, 6}, true},
		{[]int{5, 2, 6, 1, 3}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.preorder), func(t *testing.T) {
			require.Equal(t, tc.want, verifyPreorder(tc.preorder))
		})
	}
}

func verifyPreorder(preorder []int) bool {
	gt := -1
	stack := []int{}
	for i, cur := range preorder {
		if len(stack) == 0 || cur < preorder[i-1] {
			if cur <= gt {
				return false
			}
			stack = append(stack, cur)
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] < cur {
			gt = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
	}
	return true
}
