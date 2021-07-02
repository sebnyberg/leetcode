package p0689maximumsumof3nonoverlappingsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumOfThreeSubarrays(tc.nums, tc.k))
		})
	}
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {

}
