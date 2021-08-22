package p1940longestcommonsubsequencebetweensortedarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommomSubsequence(t *testing.T) {
	for _, tc := range []struct {
		arrays [][]int
		want   []int
	}{
		{[][]int{{1, 3, 4}, {1, 4, 7, 9}}, []int{1, 4}},
		{[][]int{{2, 3, 6, 8}, {1, 2, 3, 5, 6, 7, 10}, {2, 3, 4, 6, 9}}, []int{2, 3, 6}},
		{[][]int{{1, 2, 3, 4, 5}, {6, 7, 8}}, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arrays), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, longestCommomSubsequence(tc.arrays))
		})
	}
}

func longestCommomSubsequence(arrays [][]int) []int {
	// This is easier than it seems
	// The longest common sequence is simply the set of numbers which occur
	// in all arrays
	var numCount [101]int
	for _, arr := range arrays {
		for _, n := range arr {
			numCount[n]++
		}
	}
	var res []int
	wantCount := len(arrays)
	for num, count := range numCount {
		if count == wantCount {
			res = append(res, num)
		}
	}
	return res
}
