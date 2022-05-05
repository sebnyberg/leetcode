package p1999smallestgreatermultiplemadeoftwodigits

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findInteger(t *testing.T) {
	for _, tc := range []struct {
		k              int
		digit1, digit2 int
		want           int
	}{
		{2, 0, 2, 20},
		{3, 4, 2, 24},
		{2, 0, 0, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, findInteger(tc.k, tc.digit1, tc.digit2))
		})
	}
}

func findInteger(k int, digit1 int, digit2 int) int {
	// Int32 has roughly 9 digits
	// For each position there are two alternatives to choose from.
	// This gives 2^9 permutations < 1000 in total.
	// And so it is possible to try all permutations of digit1 and digit2,
	// brute-force-finding the solution
	res := dfs(0, digit1, digit2, k)
	if res == notFound {
		return -1
	}
	return res
}

const notFound = math.MaxInt32 + 1

func dfs(curVal, digit1, digit2, k int) int {
	if curVal > math.MaxInt32 {
		return notFound
	}
	res := notFound
	if curVal > 0 || digit1 > 0 {
		res = min(res, dfs((curVal*10)+digit1, digit1, digit2, k))
	}
	if curVal > 0 || digit2 > 0 {
		res = min(res, dfs((curVal*10)+digit2, digit1, digit2, k))
	}
	if curVal > k && curVal%k == 0 {
		res = min(res, curVal)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
