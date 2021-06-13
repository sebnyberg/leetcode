package p1898maxnumberofremovablecharacters

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumRemovals(t *testing.T) {
	for _, tc := range []struct {
		s         string
		p         string
		removable []int
		want      int
	}{
		{"abcacb", "ab", []int{3, 1, 0}, 2},
		{"abcbddddd", "abcd", []int{3, 2, 1, 4, 5, 6}, 1},
		{"abcab", "abc", []int{0, 1, 2, 3, 4}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			res := maximumRemovals(tc.s, tc.p, tc.removable)
			require.Equal(t, tc.want, res)
		})
	}
}

func maximumRemovals(s string, p string, removable []int) int {
	// Remove all indives from s (shift by some offset)
	bs := []byte(s)
	ps := []byte(p)
	n := len(s)
	introductionIndices := make([]int, n)
	maxK := len(removable)
	for i, removeIdx := range removable {
		introductionIndices[removeIdx] = maxK - i
	}

	// minIntroductions is the minimum number of introductions needed so that
	// s contains the pattern p
	minIntroductions := sort.Search(maxK, func(n int) bool {
		return hasSubstring(bs, ps, introductionIndices, n)
	})
	return maxK - minIntroductions
}

func hasSubstring(bs, ps []byte, introIndices []int, k int) bool {
	var j int
	n := len(ps)
	for i := range bs {
		if introIndices[i] <= k && bs[i] == ps[j] {
			j++
			if j == n {
				return true
			}
		}
	}
	return false
}
