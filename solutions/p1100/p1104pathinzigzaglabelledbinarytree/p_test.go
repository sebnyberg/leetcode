package p1104pathinzigzaglabelledbinarytree

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pathInZigZagTree(t *testing.T) {
	for i, tc := range []struct {
		label int
		want  []int
	}{
		{14, []int{1, 3, 4, 14}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, pathInZigZagTree(tc.label))
		})
	}
}

func pathInZigZagTree(label int) []int {
	var res []int

	for label >= 2 {
		res = append(res, label)
		label /= 2
		x := math.Log2(float64(label))
		sz := 1 << int(x)
		label = sz + ((sz - 1) - (label - sz))
	}

	res = append(res, 1)
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
