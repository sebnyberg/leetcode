package p1413minvaluetogetpositivestepbystepsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minStartValue(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-3, 2, -3, 4, 2}, 5},
		{[]int{1, 2}, 1},
		{[]int{1, -2, -3}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minStartValue(tc.nums))
		})
	}
}

func minStartValue(nums []int) int {
	val := nums[0]
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		val += nums[i]
		if val < minVal {
			minVal = val
		}
	}
	if minVal >= 0 {
		return 1
	}
	return -minVal + 1
}
