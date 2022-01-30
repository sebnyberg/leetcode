package p2154keepmultiplyingfoundvaluesbytwo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findFinalValue(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		original int
		want     int
	}{
		{[]int{5, 3, 6, 1, 12}, 3, 24},
		{[]int{2, 7, 9}, 4, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findFinalValue(tc.nums, tc.original))
		})
	}
}
func findFinalValue(nums []int, original int) int {
	m := map[int]struct{}{}
	for _, n := range nums {
		m[n] = struct{}{}
	}
	for {
		if _, exists := m[original]; exists {
			original *= 2
		} else {
			break
		}
	}
	return original
}
