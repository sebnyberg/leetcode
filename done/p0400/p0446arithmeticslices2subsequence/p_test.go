package p0446arithmeticslices2subsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfArithmeticSlices(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{79, 20, 64, 28, 67, 81, 60, 58, 97, 85, 92, 96, 82, 89, 46, 50, 15, 2, 36, 44, 54, 2, 90, 37, 7, 79, 26, 40, 34, 67, 64, 28, 60, 89, 46, 31, 9, 95, 43, 19, 47, 64, 48, 95, 80, 31, 47, 19, 72, 99, 28, 46, 13, 9, 64, 4, 68, 74, 50, 28, 69, 94, 93, 3, 80, 78, 23, 80, 43, 49, 77, 18, 68, 28, 13, 61, 34, 44, 80, 70, 55, 85, 0, 37, 93, 40, 47, 47, 45, 23, 26, 74, 45, 67, 34, 20, 33, 71, 48, 96}, 1030},
		{[]int{2, 4, 6, 8, 10}, 7},
		{[]int{7, 7, 7, 7, 7}, 16},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfArithmeticSlices(tc.nums))
		})
	}
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	seqs := make([]map[int]int, n)
	for i := range seqs {
		seqs[i] = make(map[int]int)
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			res += seqs[j][d]
			seqs[i][d] += seqs[j][d] + 1
		}
	}
	return res
}
