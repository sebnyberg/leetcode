package p1995countspecialquadruplets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countQuadruplets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 6}, 1},
		{[]int{3, 3, 6, 4, 5}, 0},
		{[]int{1, 1, 1, 3, 5}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countQuadruplets(tc.nums))
		})
	}
}

func countQuadruplets(nums []int) int {
	// Ugly but whatever
	var res int
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						res++
					}
				}
			}
		}
	}
	return res
}
