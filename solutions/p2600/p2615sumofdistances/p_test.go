package p2615sumofdistances

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distance(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want []int64
	}{
		{[]int{1, 3, 1, 1, 2}, []int64{5, 0, 3, 4, 0}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, distance(tc.nums))
		})
	}
}

func distance(nums []int) []int64 {
	n := len(nums)

	count := make(map[int]int)
	sum := make(map[int]int)
	last := make(map[int]int)

	res := make([]int64, n)
	for i := range nums {
		sum[nums[i]] += abs(i-last[nums[i]]) * count[nums[i]]
		res[i] += int64(sum[nums[i]])
		count[nums[i]]++
		last[nums[i]] = i
	}
	for k := range count {
		delete(count, k)
		delete(sum, k)
	}
	for i := n - 1; i >= 0; i-- {
		sum[nums[i]] += abs(i-last[nums[i]]) * count[nums[i]]
		res[i] += int64(sum[nums[i]])
		count[nums[i]]++
		last[nums[i]] = i
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
