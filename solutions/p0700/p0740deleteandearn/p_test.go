package p0740deleteandearn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteAndEarn(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, deleteAndEarn(tc.nums))
		})
	}
}

func deleteAndEarn(nums []int) int {
	var count [10001]uint16
	for _, num := range nums {
		count[num]++
	}
	var maxPrevDeleted, maxPrevNotDeleted int
	for x, n := range count {
		maxPrevDeleted, maxPrevNotDeleted = max(maxPrevDeleted, maxPrevNotDeleted+int(n)*x),
			maxPrevDeleted
	}
	return max(maxPrevDeleted, maxPrevNotDeleted)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
