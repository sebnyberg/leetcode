package p1764formarraybyconcatenatingsubarraysofanotherarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canChoose(t *testing.T) {
	for _, tc := range []struct {
		groups [][]int
		nums   []int
		want   bool
	}{
		{[][]int{{6551094, 9427527, 2052462, 3481286, -7620442}, {8495362, -1820796}, {-1005271, -6911519}, {-9667242, 9997184, -9316362}, {-9278108, -7479063, -7573091, -1775876, -2612810, -241649}},
			[]int{6551094, 6551094, 9427527, 2052462, 3481286, -7620442, -7620442, 8495362, -1820796, -1005271, -6911519, -9667242, 9997184, -9316362, 9997184, -9278108, -7479063, -7573091, -1775876, -2612810, -241649}, true},
		{[][]int{{21, 22, 21, 22, 21, 30}}, []int{21, 22, 21, 22, 21, 22, 21, 30}, true},
		{[][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}, true},
		{[][]int{{10, -2}, {1, 2, 3, 4}}, []int{1, 2, 3, 4, 10, -2}, false},
		{[][]int{{1, 2, 3}, {3, 4}}, []int{7, 7, 1, 2, 3, 4, 7, 7}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.groups), func(t *testing.T) {
			require.Equal(t, tc.want, canChoose(tc.groups, tc.nums))
		})
	}
}

func canChoose(groups [][]int, nums []int) bool {
	var pos int
	lps := make([]uint16, 0, len(nums))
	for _, group := range groups {
		// Construct LPS
		n := len(group)
		lps = lps[:1]
		for i := 1; i < n; i++ {
			j := lps[i-1]
			for j > 0 && group[i] != group[j] {
				j = lps[j-1]
			}
			if group[i] == group[j] {
				j++
			}
			lps = append(lps, j)
		}

		// Match
		var i int
		for pos < len(nums) && i < n {
			if nums[pos] == group[i] {
				i++
			} else {
				for i > 0 && nums[pos] != group[i] {
					i = int(lps[i-1])
				}
				if i > 0 || nums[pos] == group[i] {
					i++
				}
			}
			pos++
		}
		if i != n {
			return false
		}
	}
	return true
}
