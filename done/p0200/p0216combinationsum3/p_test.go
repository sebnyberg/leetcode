package p0216combinationsum3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum3(t *testing.T) {
	for _, tc := range []struct {
		k    int
		n    int
		want [][]int
	}{
		// {3, 7, [][]int{{1, 2, 4}}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.k, tc.n), func(t *testing.T) {
			require.Contains(t, tc.want, combinationSum3(tc.k, tc.n))
		})
	}
}

func combinationSum3(k int, n int) [][]int {
	var s sumFinder
	s.k = k
	s.findSums(1, n, []int{})
	return s.results
}

type sumFinder struct {
	k       int
	results [][]int
}

func (s *sumFinder) findSums(cur int, n int, prefix []int) {
	if len(prefix) == s.k {
		if n != 0 {
			return
		}
		prefixCpy := make([]int, len(prefix))
		copy(prefixCpy, prefix)
		s.results = append(s.results, prefixCpy)
		return
	}
	if n < 0 || cur > 9 {
		return
	}

	// continue without...
	s.findSums(cur+1, n, prefix)

	// ...and with cur
	prefixCpy := make([]int, len(prefix)+1)
	copy(prefixCpy, prefix)
	prefixCpy[len(prefixCpy)-1] = cur
	s.findSums(cur+1, n-cur, prefixCpy)
}
