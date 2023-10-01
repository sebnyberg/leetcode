package p2763sumofimbalancenumbersofallsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumImbalanceNumbers(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 4}, 3},
		{[]int{1, 3, 3, 3, 5}, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, sumImbalanceNumbers(tc.nums))
		})
	}
}

func sumImbalanceNumbers(nums []int) int {
	// It may actually be possible to just brute-force this. It's worth a try
	seen := make(map[int]bool)
	var res int
	for i := 0; i < len(nums)-1; i++ {
		for k := range seen {
			delete(seen, k)
		}
		seen[nums[i]] = true

		var imbalance int
		for j := i + 1; j < len(nums); j++ {
			// Let's assume that each number is expected to increase the total
			// imbalance by 1.
			imbalance++

			// This is not true if we've already seen the number
			if seen[nums[j]] {
				imbalance--
			} else {
				seen[nums[j]] = true

				// Or if this number is too close to its predecessor
				if seen[nums[j]-1] {
					imbalance--
				}
				// Or if this number is too close to its successor
				if seen[nums[j]+1] {
					imbalance--
				}
			}

			res += imbalance
		}
	}
	return res
}
