package p0633sumofsquarenumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_judgeSquareSum(t *testing.T) {
	for _, tc := range []struct {
		c    int
		want bool
	}{
		{5, true},
		{3, false},
		{4, true},
		{2, true},
		{1, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.c), func(t *testing.T) {
			require.Equal(t, tc.want, judgeSquareSum(tc.c))
		})
	}
}

// judgeSquareSum returns true if there exists two integers a and b such that
// a^2 + b^2 = c
func judgeSquareSum(c int) bool {
	seen := make(map[int]struct{})
	for i := 0; i*i <= c; i++ {
		val := i * i
		seen[i*i] = struct{}{}
		if _, exists := seen[c-val]; exists {
			return true
		}
	}
	return false
}

// func judgeSquareSum2(c int) bool {
// 	l, r := 0, newtonRaphsonSqrt(c)

// }

// func newtonRaphsonSqrt(x int) int {
// 	c := x
// 	for x*x > c {
// 		x = (x + (c / x)) / 2
// 	}
// 	return x
// }
