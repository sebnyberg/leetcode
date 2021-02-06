package p0045jumpgame2

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jump(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, jump(tc.nums))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	n := len(nums)
	minjumps := make([]int, n)
	for i := range minjumps {
		minjumps[i] = math.MaxInt64
	}
	minjumps[0] = 0
	for i := range nums {
		jumplen := nums[i]
		minjump := minjumps[i]
		for j := 1; i+j < n && j <= jumplen; j++ {
			minjumps[i+j] = min(minjumps[i+j], minjump+1)
		}
	}

	return minjumps[len(nums)-1]
}
