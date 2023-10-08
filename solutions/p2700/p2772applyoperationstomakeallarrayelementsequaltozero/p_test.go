package p2772applyoperationstomakeallarrayelementsequaltozero

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkArray(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{60, 72, 87, 89, 63, 52, 64, 62, 31, 37, 57, 83, 98, 94, 92, 77, 94, 91, 87, 100, 91, 91, 50, 26}, 4, true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, checkArray(tc.nums, tc.k))
		})
	}
}

func checkArray(nums []int, k int) bool {
	n := len(nums)
	deltas := make([]int, n+1)
	var delta int
	for i := range nums {
		delta += deltas[i]
		nums[i] -= delta
		delta += nums[i]
		if nums[i] < 0 {
			return false
		}
		if nums[i] == 0 {
			continue
		}
		if i+k > n {
			return false
		}
		deltas[i+k] -= nums[i]
	}
	return true
}
