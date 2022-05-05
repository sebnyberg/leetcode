package p0414thirdmaximumnumber

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_thirdMax(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, thirdMax(tc.nums))
		})
	}
}

func thirdMax(nums []int) int {
	unset := math.MinInt32 - 1
	var max [3]int
	for i := range max {
		max[i] = unset
	}
	for _, num := range nums {
		for i := range max {
			if num == max[i] {
				break
			}
			if num > max[i] {
				max[i], num = num, max[i]
			}
		}
	}
	if max[2] == unset {
		return max[0]
	}
	return max[2]
}
