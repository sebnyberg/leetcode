package p3208alternatinggroupsii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfAlternatingGroups(t *testing.T) {
	for _, tc := range []struct {
		colors []int
		k      int
		want   int
	}{
		{[]int{0, 1, 0, 1, 0}, 3, 3},
		{[]int{0, 1, 0, 0, 1, 0, 1}, 6, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.colors), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfAlternatingGroups(tc.colors, tc.k))
		})
	}
}

func numberOfAlternatingGroups(colors []int, k int) int {
	var numAlternating int
	for i := 1; i < k-1; i++ {
		if colors[i] != colors[i-1] {
			numAlternating++
		}
	}
	n := len(colors)
	var res int
	for i := 0; i < n; i++ {
		// add last alternating color (if any)
		a, b := colors[(i+k-2)%n], colors[(i+k-1)%n]
		if a != b {
			numAlternating++
		}
		if numAlternating == k-1 {
			res++
		}
		if colors[(i+1)%n] != colors[i] {
			numAlternating--
		}
	}
	return res
}
