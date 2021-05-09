package p0312burstbaloons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCoins(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxCoins(tc.nums))
		})
	}
}

func maxCoins(nums []int) int {
	return 0
}
