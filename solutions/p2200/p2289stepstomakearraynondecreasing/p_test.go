package p2289stepstomakearraynondecreasing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalSteps(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{82, 228, 317, 559, 670, 1535, 1977, 890, 1243, 1502, 1720, 1808, 1819, 1966, 105, 170, 194, 320, 385, 433, 607, 633, 1144, 1195, 1365, 1490, 1778, 921, 1560, 6, 68, 69, 100, 341, 595, 679, 725, 775, 887, 1316, 296, 613, 658, 682, 777, 1203, 1350, 1431, 161, 374, 784, 794, 863, 1080, 1149, 1509, 1525, 128, 437, 996, 1045, 1061, 1102, 1238, 1624, 1706, 1961, 808, 950, 1166, 1531, 1537, 1732, 866, 1279, 1494, 1527, 1595}, 18},
		{[]int{7, 15, 5, 5, 6, 10, 7, 13}, 5},
		{[]int{5, 14, 15, 2, 11, 5, 13, 15}, 3},
		{[]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}, 6},
		{[]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}, 3},
		{[]int{4, 5, 7, 7, 13}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, totalSteps(tc.nums))
		})
	}
}

func totalSteps(nums []int) int {
	n := len(nums)
	const idx = 0
	const count = 1
	stack := [][2]int{}
	var res int
	for i := n - 1; i >= 0; i-- {
		var c int
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1][idx]] {
			c = max(c+1, stack[len(stack)-1][count])
			stack = stack[:len(stack)-1]
			res = max(res, c)
		}
		stack = append(stack, [2]int{i, c})
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
