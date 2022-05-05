package p0481magicalstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_magicalString(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{20, 3},
		{6, 3},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, magicalString(tc.n))
		})
	}
}

func magicalString(n int) int {
	// The starting point is '1'
	// Then, since the group must hold 1 number, the next number must be 2
	stack := make([]int, 0, n)
	stack = append(stack, 1, 2, 2)
	var i, j int = 2, 3
	// 1221121221221121122
	for j < n {
		// remove first element from queue
		if stack[i] == 1 {
			j++
			// End must add a single element of another kind than the current end
			if stack[len(stack)-1] == 1 {
				stack = append(stack, 2)
			} else {
				stack = append(stack, 1)
			}
		} else {
			j += 2
			// End must add two elements of another kind than the current end
			if stack[len(stack)-1] == 1 {
				stack = append(stack, 2, 2)
			} else {
				stack = append(stack, 1, 1)
			}
		}
		i++
	}
	var res int
	for i := 0; i < n; i++ {
		if stack[i] == 1 {
			res++
		}
	}
	return res
}
