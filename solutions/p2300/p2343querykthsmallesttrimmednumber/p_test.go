package p2343querykthsmallesttrimmednumber

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestTrimmedNumbers(t *testing.T) {
	for _, tc := range []struct {
		nums    []string
		queries [][]int
		want    []int
	}{
		{
			[]string{"102", "473", "251", "814"},
			[][]int{{1, 1}, {2, 3}, {4, 2}, {1, 2}},
			[]int{2, 2, 1, 0},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, smallestTrimmedNumbers(tc.nums, tc.queries))
		})
	}
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	// 100 x 100 can be brute-forced...
	// We can create a list which carries the original position as information.
	// On query, we sort according to the criteria and pick kth element,
	// then we unsort using the original position of the string
	type strAndPos struct {
		pos int
		s   string
	}
	n := len(nums)
	strs := make([]strAndPos, n)
	for i, x := range nums {
		strs[i].pos = i
		strs[i].s = x
	}
	reset := func() {
		sort.Slice(strs, func(i, j int) bool {
			return strs[i].pos < strs[j].pos
		})
	}
	m := len(strs[0].s)
	res := make([]int, len(queries))
	for i, q := range queries {
		k, trim := q[0], q[1]
		sort.Slice(strs, func(i, j int) bool {
			s1 := strs[i].s[m-trim:]
			s2 := strs[j].s[m-trim:]
			if s1 == s2 {
				return strs[i].pos < strs[j].pos
			}
			return s1 < s2
		})
		res[i] = strs[k-1].pos
		reset()
	}
	return res
}
