package p1531stringcompressionii

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getLengthOfOptimalCompression(t *testing.T) {
	for i, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"aaabcccd", 2, 4},
		{"aabbaa", 2, 2},
		{"aaaaaaaaaaaaaa", 0, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getLengthOfOptimalCompression(tc.s, tc.k))
		})
	}
}

func getLengthOfOptimalCompression(s string, k int) int {
	m := make(map[key]int)
	res := dp(m, s, 0, k, 0, 0)
	return res
}

type key struct {
	i         uint8
	k         uint8
	prevCount uint8
	prev      uint8
}

func dp(m map[key]int, s string, i, k, prevCount int, prev byte) int {
	if i == len(s) {
		return 0
	}
	kk := key{
		i:         uint8(i),
		k:         uint8(k),
		prevCount: uint8(prevCount),
		prev:      uint8(prev),
	}
	if v, exists := m[kk]; exists {
		return v
	}

	// let's consider a skip first.
	res := math.MaxInt32
	if k > 0 {
		res = min(res, dp(m, s, i+1, k-1, prevCount, prev))
	}

	// then a merge
	var cost int
	newCount := prevCount + 1
	if s[i] != prev {
		newCount = 1
	}
	if newCount <= 2 || newCount == 10 || newCount == 100 {
		cost++
	}
	res = min(res, cost+dp(m, s, i+1, k, newCount, s[i]))
	m[kk] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
