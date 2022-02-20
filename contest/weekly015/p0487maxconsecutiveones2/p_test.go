package p0487maxconsecutiveones2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 1, 1, 1}, 5},
		{[]int{1, 1, 1, 1, 0}, 5},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 0, 1, 1, 0}, 4},
		{[]int{1, 0, 1, 1, 0, 1}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxConsecutiveOnes(tc.nums))
		})
	}
}

func findMaxConsecutiveOnes(nums []int) int {
	nums = append(nums, 0)
	pos := [3]int{-1, -1, -1}
	var res int
	for i, num := range nums {
		if num == 1 {
			continue
		}
		pos[0], pos[1], pos[2] = pos[1], pos[2], i
		if d := pos[2] - pos[0] - 1; d > res {
			res = d
		}
	}
	return res
}
