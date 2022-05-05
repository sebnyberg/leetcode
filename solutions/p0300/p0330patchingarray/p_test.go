package p0330patchingarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPatches(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		n    int
		want int
	}{
		{[]int{1, 3}, 6, 1},
		{[]int{1, 5, 10}, 20, 2},
		{[]int{1, 2, 2}, 5, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minPatches(tc.nums, tc.n))
		})
	}
}

func minPatches(nums []int, n int) int {
	want := 1
	var patches int
	for i := 0; i < len(nums); {
		if nums[i] <= want {
			want += nums[i]
			i++
		} else {
			want *= 2
			patches++
		}
		if want > n {
			return patches
		}
	}
	for want <= n {
		want *= 2
		patches++
	}
	return patches
}
