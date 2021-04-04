package p0274hindex

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hIndex(t *testing.T) {
	for _, tc := range []struct {
		citations []int
		want      int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.citations), func(t *testing.T) {
			require.Equal(t, tc.want, hIndex(tc.citations))
		})
	}
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	for i, citation := range citations {
		if citation < i+1 {
			return i
		}
	}
	return len(citations)
}
