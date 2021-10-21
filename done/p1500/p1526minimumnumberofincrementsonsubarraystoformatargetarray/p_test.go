package p1526minimumnumberofincrementsonsubarraystoformatargetarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minNumberOperations(t *testing.T) {
	for _, tc := range []struct {
		target []int
		want   int
	}{
		{[]int{1, 2, 3, 2, 1}, 3},
		{[]int{3, 1, 1, 2}, 4},
		{[]int{3, 1, 5, 4, 2}, 7},
		{[]int{1, 1, 1, 1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, minNumberOperations(tc.target))
		})
	}
}

func minNumberOperations(target []int) int {
	operations := target[0]
	reusable := target[0]
	n := len(target)
	for i := 1; i < n; i++ {
		if target[i] < reusable {
			reusable = target[i]
		} else if target[i] > reusable {
			operations += target[i] - reusable
			reusable = target[i]
		}
	}
	return operations
}
