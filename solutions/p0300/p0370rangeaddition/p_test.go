package p0370rangeaddition

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getModifiedArray(t *testing.T) {
	for _, tc := range []struct {
		length  int
		updates [][]int
		want    []int
	}{
		{5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}, []int{-2, 0, 3, 5, 3}},
		{10, [][]int{{2, 4, 6}, {5, 6, 8}, {1, 9, -4}}, []int{0, -4, 2, 2, 2, 4, 4, -4, -4, -4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.length), func(t *testing.T) {
			require.Equal(t, tc.want, getModifiedArray(tc.length, tc.updates))
		})
	}
}

func getModifiedArray(length int, updates [][]int) []int {
	// An update either increases or decreases the value at a position and
	// following positions. Iterate over updates and put these changes into
	// the ops slice:
	ops := make([]int, length+1)
	for _, update := range updates {
		ops[update[0]] += update[2]
		ops[update[1]+1] -= update[2]
	}

	result := make([]int, length)
	result[0] = ops[0]
	for i := 1; i < len(result); i++ {
		result[i] = result[i-1] + ops[i]
	}
	return result
}
