package p1224maximumequalfrequency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxEqualFreq(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1, 1, 5, 3, 3, 5}, 7},
		{[]int{2, 2, 1, 1, 5, 3, 3, 5}, 7},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxEqualFreq(tc.nums))
		})
	}
}

func maxEqualFreq(nums []int) int {
	freqCount := make(map[int]int)
	count := make(map[int]int)
	res := 1
	for i, x := range nums {
		c := count[x]
		if c != 0 {
			freqCount[c]--
			if freqCount[c] == 0 {
				delete(freqCount, c)
			}
		}
		count[x]++
		freqCount[count[x]]++

		// a prefix is valid iff there are at most two distinct frequencies, and
		// the count of values in each differs by at most 1.
		if len(freqCount) > 2 {
			continue
		}

		var ks []int
		var vs []int
		for k, v := range freqCount {
			ks = append(ks, k)
			vs = append(vs, v)
		}

		if len(freqCount) == 1 {
			// freq must be 1
			for k, cnt := range freqCount {
				// if the freq is 1, just remove any of the numbers
				if k == 1 {
					res = max(res, i+1)
				}
				// if the count of distinct numbers is 1, remove one of the
				// numbers from that group of same numbers
				if cnt == 1 {
					res = max(res, i+1)
				}
			}
			continue
		}

		// There are two cases: either the two values are next to each other,
		// and the upper count is 1, or the lower frequency is 1 and the count
		// is 1. In both cases, either the first or second count must be 1
		if vs[0] != 1 && vs[1] != 1 {
			continue
		}
		if vs[0] == 1 && (ks[0] == 1 || ks[0]-ks[1] == 1) {
			res = max(res, i+1)
			continue
		}
		if vs[1] == 1 && (ks[1] == 1 || ks[1]-ks[0] == 1) {
			res = max(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
