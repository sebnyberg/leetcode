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
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.citations), func(t *testing.T) {
			require.Equal(t, tc.want, hIndex(tc.citations))
		})
	}
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	_ = n
	// mid := n / 2
	return 0
}
