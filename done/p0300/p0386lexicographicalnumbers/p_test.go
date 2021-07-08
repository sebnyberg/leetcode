package p0386lexicographicalnumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lexicalOrder(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want []int
	}{
		{99, []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 4, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 5, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 6, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 7, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 8, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 9, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}},
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{2, []int{1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, lexicalOrder(tc.n))
		})
	}
}

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	stack := make(intStack, 0, 10)
	// x contains the number to increment
	// x must not end with 9 and be < n
	var x int
	for {
		x++
		res = append(res, x)
		for x*10 <= n {
			stack.push(x) // Return to this number later on
			x *= 10
			res = append(res, x)
		}

		// Find next number to increment from
		// Any number which would go beyond 'n' or ends with 9 is not a valid number
		for x >= n || x%10 == 9 {
			if len(stack) == 0 {
				return res
			}
			x = stack.pop()
		}
	}
}

type intStack []int

func (s *intStack) push(x int) { *s = append(*s, x) }
func (s *intStack) pop() int {
	n := len(*s)
	el := (*s)[n-1]
	*s = (*s)[:n-1]
	return el
}
