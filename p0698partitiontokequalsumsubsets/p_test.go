package p0698partitiontokequalsumsubsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPartitionKSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{4, 3, 2, 3, 5, 2, 1}, 4, true},
		{[]int{1, 2, 3, 4}, 3, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, canPartitionKSubsets(tc.nums, tc.k))
		})
	}
}

func canPartitionKSubsets(nums []int, k int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	want := sum / k
}
