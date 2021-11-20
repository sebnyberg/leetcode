package p0540singleelementinasortedarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNonDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, singleNonDuplicate(tc.nums))
		})
	}
}

func singleNonDuplicate(nums []int) int {
	idx := sort.Search((len(nums)-1)/2, func(i int) bool {
		return nums[i] == nums[i+1]
	})
	return nums[idx]
}
