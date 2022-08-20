package p0830positionsoflargegroups

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largeGroupPositions(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want [][]int
	}{
		{"abbxxxxzzy", [][]int{{3, 6}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, largeGroupPositions(tc.s))
		})
	}
}

func largeGroupPositions(s string) [][]int {
	var i int
	res := make([][]int, 0)
	for j := range s {
		if s[j] == s[i] {
			continue
		}
		if j-i >= 3 {
			res = append(res, []int{i, j - 1})
		}
		i = j
	}
	if len(s)-i >= 3 {
		res = append(res, []int{i, len(s) - 1})
	}
	return res
}
