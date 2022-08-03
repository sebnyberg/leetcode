package p0598rangeadditionii

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCount(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{3, 3, leetcode.ParseMatrix("[[2,2],[3,3]]"), 4},
		{3, 3, leetcode.ParseMatrix("[[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3]]"), 4},
		{3, 3, leetcode.ParseMatrix("[[]]"), 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, maxCount(tc.m, tc.n, tc.ops))
		})
	}
}

func maxCount(m int, n int, ops [][]int) int {
	for _, o := range ops {
		m = min(m, o[0])
		n = min(n, o[1])
	}
	return m * n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
