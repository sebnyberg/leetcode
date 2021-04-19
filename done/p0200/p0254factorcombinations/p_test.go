package p0254factorcombinations

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getFactors(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want [][]int
	}{
		{12, [][]int{{2, 6}, {3, 4}, {2, 3, 2}}},
		{1, [][]int{}},
		{37, [][]int{}},
		{32, [][]int{{2, 16}, {4, 8}, {2, 8, 2}, {4, 4, 2}, {2, 4, 2, 2}, {2, 2, 2, 2, 2}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			result := getFactors(tc.n)
			require.ElementsMatch(t, tc.want, result)
		})
	}
}

func getFactors(n int) [][]int {
	if n == 1 {
		return [][]int{}
	}
	resp := findFactors(n, 2)
	return resp
}

func findFactors(n int, start int) [][]int {
	res := make([][]int, 0)
	for a := start; a <= int(math.Sqrt(float64(n))); a++ {
		if n%a == 0 {
			res = append(res, []int{a, n / a})
			for _, subFactors := range findFactors(n/a, a) {
				res = append(res, append(subFactors, a))
			}
		}
	}
	return res
}
