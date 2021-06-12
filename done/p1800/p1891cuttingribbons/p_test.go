package p1891cuttingribbons

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxLength(t *testing.T) {
	for _, tc := range []struct {
		ribbons []int
		k       int
		want    int
	}{
		{[]int{9, 7, 5}, 3, 5},
		{[]int{7, 5, 9}, 4, 4},
		{[]int{5, 7, 9}, 22, 0},
		{[]int{100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 1, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000, 100000},
			49, 100000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ribbons), func(t *testing.T) {
			require.Equal(t, tc.want, maxLength(tc.ribbons, tc.k))
		})
	}
}

func maxLength(ribbons []int, k int) int {
	var maxLength int
	for _, r := range ribbons {
		if r > maxLength {
			maxLength = r
		}
	}
	l, r := 0, maxLength
	var got int
	for l < r {
		m := (l + r + 1) / 2
		// calculate number of ribbons
		got = 0
		for _, ribbon := range ribbons {
			got += ribbon / m
			if got >= k {
				break
			}
		}
		if got < k {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func builtinSearch(ribbons []int, k int) int {
	var maxLength int
	for _, r := range ribbons {
		maxLength = max(maxLength, r)
	}
	res := sort.Search(maxLength+1, func(m int) bool {
		got := 0
		if m == 0 {
			return false
		}
		for _, ribbon := range ribbons {
			got += ribbon / m
			if got >= k {
				return false
			}
		}
		return true
	})
	return res - 1
}
