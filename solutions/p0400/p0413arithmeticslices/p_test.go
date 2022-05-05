package p0413arithmeticslices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfArithmeticSlices(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		want int
	}{
		{[]int{1, 3, 5, 7, 9}, 6},
		{[]int{7, 7, 7, 7}, 3},
		{[]int{3, -1, -5, -9}, 3},
		{[]int{1, 2, 3, 4}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfArithmeticSlices(tc.A))
		})
	}
}

func numberOfArithmeticSlices(nums []int) int {
	count := 2
	var res int
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			count++
		} else {
			count = 2
		}
		res += count - 2
	}
	return res
}
