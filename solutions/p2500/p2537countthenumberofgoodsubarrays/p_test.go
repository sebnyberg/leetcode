package p2537countthenumberofgoodsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countGood(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{3, 1, 4, 3, 2, 2, 4}, 2, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countGood(tc.nums, tc.k))
		})
	}
}

func countGood(nums []int, k int) int64 {
	// Maintain a window which is not good
	// Any subarray that is longer than that window is good by definition
	//
	var j int
	count := make(map[int]int)
	var npairs int
	var res int64
	for _, x := range nums {
		npairs += count[x]
		count[x]++
		for npairs >= k {
			count[nums[j]]--
			npairs -= count[nums[j]]
			j++
		}
		res += int64(j)
	}
	return res
}
