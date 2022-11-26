package p1006clumsyfactorial

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_clumsy(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		// {4, 7},
		{10, 12},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, clumsy(tc.n))
		})
	}
}

func clumsy(n int) int {
	stack := []int{n}
	n--
	for i := 0; n > 0; i++ {
		switch i % 4 {
		case 0:
			stack[len(stack)-1] *= n
		case 1:
			stack[len(stack)-1] /= n
		case 2:
			stack = append(stack, n)
		case 3:
			stack = append(stack, -n)
		}
		n--
	}
	res := stack[0]
	for _, x := range stack[1:] {
		res += x
	}
	return res
}
