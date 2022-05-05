package p0169majorityelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, majorityElement(tc.nums))
		})
	}
}

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	want := len(nums) / 2
	for _, num := range nums {
		if counts[num] == want {
			return num
		}
		counts[num]++
	}
	return 0
}
