package p0089graycode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_grayCode(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 3, 2}},
		{1, []int{0, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, grayCode(tc.n))
		})
	}
}

func grayCode(n int) []int {
	// Neat solution found in discussion
	nresults := 1 << n
	res := make([]int, nresults)
	for i := 0; i < nresults; i++ {
		res[i] = i ^ (i >> 1)
	}
	return res
}

func grayCodeFirst(n int) []int {
	// First attempt done without checking Google
	nresults := 1 << n
	res := make([]int, nresults)
	for i := 0; i < nresults; i++ {
		var k int
		for j := 0; j < n; j++ {
			k |= (((i + 1<<j) / (1 << (j + 1))) % 2) << j
		}
		res[i] = k
	}
	return res
}
