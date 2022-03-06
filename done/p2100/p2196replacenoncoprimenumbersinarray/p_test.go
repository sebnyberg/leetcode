package p2196replacenoncoprimenumbersinarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_replaceNonCoprimes(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{13, 13, 13, 13, 13, 13, 13, 13}, []int{13}},
		{[]int{287, 41, 49, 287, 899, 23, 23, 20677, 5, 825}, []int{2009, 20677, 825}},
		{[]int{6, 4, 3, 2, 7, 6, 2}, []int{12, 7, 6}},
		{[]int{2, 2, 1, 1, 3, 3, 3}, []int{2, 1, 1, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, replaceNonCoprimes(tc.nums))
		})
	}
}

func replaceNonCoprimes(nums []int) []int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	lcm := func(a, b int) int {
		return a * b / gcd(a, b)
	}

	stack := make([]int, 0, len(nums))
	var n int

	for _, num := range nums {
		stack = append(stack, num)
		n++
		for n > 1 && gcd(stack[n-2], stack[n-1]) != 1 {
			stack[n-2] = lcm(stack[n-2], stack[n-1])
			stack = stack[:n-1]
			n--
		}
	}
	return stack
}
