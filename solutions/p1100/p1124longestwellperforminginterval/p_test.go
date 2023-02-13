package p1124longestwellperforminginterval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestWPI(t *testing.T) {
	for i, tc := range []struct {
		hours []int
		want  int
	}{
		{[]int{8, 10, 6, 16, 5}, 3},
		{[]int{9, 6, 9}, 3},
		{[]int{6, 9, 9}, 3},
		{[]int{6, 6, 6}, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestWPI(tc.hours))
		})
	}
}

func longestWPI(hours []int) int {
	// We can keep track of the current tiring/non-tiring delta for a given
	// position, and the earliest instance of each delta.
	//
	// At a given delta, if we can deduct a sequence such that the current delta
	// is at least one, then that sequence is valid.
	//
	// For example, if the current delta is -2, and we need a delta of 1, then
	// we must remove -3.
	m := make(map[int]int)
	m[0] = -1
	var delta int
	var res int
	for i := range hours {
		if hours[i] <= 8 {
			delta--
		} else {
			delta++
		}
		if delta > 0 {
			res = max(res, i+1)
		} else {
			want := delta - 1
			if j, exists := m[want]; exists {
				res = max(res, i-j)
			}
		}
		if _, exists := m[delta]; !exists {
			m[delta] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
