package p2171removingminimumnumberofmagicbeans

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumRemoval(t *testing.T) {
	for _, tc := range []struct {
		beans []int
		want  int64
	}{
		{[]int{4, 1, 6, 5}, 4},
		{[]int{2, 10, 3, 2}, 7},
		{[]int{2, 2, 2, 2, 2}, 0},
		{[]int{1, 1, 2, 2, 2, 2}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.beans), func(t *testing.T) {
			require.Equal(t, tc.want, minimumRemoval(tc.beans))
		})
	}
}

func minimumRemoval(beans []int) int64 {
	n := len(beans)

	// Sort beans by size
	sort.Ints(beans)
	presum := make([]int, n+1)
	for i, n := range beans {
		presum[i+1] = presum[i] + n
	}
	beans = append(beans, math.MaxInt64) // sentinel

	// For each bean value,
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		// Remove all values prior to this position, and set all values after this
		// position to be equal to the current number
		v := presum[i] + (presum[n] - presum[i]) - (n-i)*beans[i]
		if v < res {
			res = v
		}

		// Forward so that next value is unique
		for i < len(beans)-1 && beans[i+1] == beans[i] {
			i++
		}
	}
	return int64(res)
}
