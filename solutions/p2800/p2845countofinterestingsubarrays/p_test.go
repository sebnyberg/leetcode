package p2845countofinterestingsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countInterestingSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		modulo int
		k      int
		want   int64
	}{
		{[]int{3, 2, 4}, 2, 1, 3},
		{[]int{3, 1, 9, 6}, 3, 0, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countInterestingSubarrays(tc.nums, tc.modulo, tc.k))
		})
	}
}

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	cnt := map[int]int{0: 1}
	prefix, res := 0, int64(0)
	for _, num := range nums {
		if num%modulo == k {
			prefix++
		}
		key := (prefix - k + modulo) % modulo
		res += int64(cnt[key])
		cnt[prefix%modulo]++
	}
	return res
}
