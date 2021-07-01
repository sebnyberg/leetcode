package p0898bitwiseorofsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subarrayBitwiseORs(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{0}, 1},
		{[]int{1, 1, 2}, 3},
		{[]int{1, 2, 4}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, subarrayBitwiseORs(tc.arr))
		})
	}
}

func subarrayBitwiseORs(arr []int) int {
	prev := []uint32{0}
	cur := []uint32{0}
	seen := make(map[uint32]bool)
	for _, num := range arr {
		cur = cur[:1]
		for _, prevNum := range prev {
			new := prevNum | uint32(num)
			seen[new] = true
			if cur[len(cur)-1] != new {
				cur = append(cur, new)
			}
		}
		cur, prev = prev, cur
	}
	return len(seen)
}
