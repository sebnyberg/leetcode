package p1920buildarrayfrompermutation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, buildArray(tc.nums))
		})
	}
}

func buildArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = nums[nums[i]]
	}
	return res
}
