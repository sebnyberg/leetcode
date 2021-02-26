package p0946validatestackseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateStackSequences(t *testing.T) {
	for _, tc := range []struct {
		pushed []int
		popped []int
		want   bool
	}{
		{[]int{1, 0}, []int{1, 0}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.pushed), func(t *testing.T) {
			require.Equal(t, tc.want, validateStackSequences(tc.pushed, tc.popped))
		})
	}

}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}

	stack := make([]int, 0, len(pushed))
	stack = append(stack, pushed[0])

	var i, j int
	stacklen := 1
	for {
		for stacklen == 0 || stack[stacklen-1] != popped[j] {
			i++
			if i == len(pushed) {
				return false
			}
			stack = append(stack, pushed[i])
			stacklen++
		}

		// pop from stack and move popped index forward
		stack = stack[:stacklen-1]
		stacklen--
		j++
		if j == len(popped) {
			return true
		}
	}
}
