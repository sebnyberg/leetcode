package p0477totalhammingdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalHammingDistance(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{4, 14, 2}, 6},
		{[]int{4, 14, 4}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, totalHammingDistance(tc.nums))
		})
	}
}

func totalHammingDistance(nums []int) int {
	n := len(nums)
	var res int
	for i := 0; i < 32; i++ {
		var k int
		for _, num := range nums {
			if num&(1<<i) > 0 {
				k++
			}
		}
		res += (n - k) * k
	}
	return res
}
