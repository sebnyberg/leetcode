package p2588countthenumberofbeautifulsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_beautifulSubarrays(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{4, 3, 1, 2, 4}, 2},
		{[]int{1, 10, 4}, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, beautifulSubarrays(tc.nums))
		})
	}
}

func beautifulSubarrays(nums []int) int64 {
	// XOR will cancel out duplicate bits
	// Cancellation order does not matter for a subarray of numbers. Only
	// whether bits could be canceled or not.
	//
	// When considering a number in nums, any subarray that ends in the previous
	// position and has an xor value equal to num can be combined to get a
	// subarray with value 0. If the current running xor is known, then it is
	// possible to find such a subarray by finding a subarray from the start
	// that contains what would need to remove.
	//
	// Sorry bout the dense explanation here, it is effectively the same
	// approach as for a subarray sum but for XOR instead.
	//
	count := make(map[int]int)
	count[0] = 1
	var xor int
	var res int
	for _, x := range nums {
		want := x ^ xor
		res += count[want]
		xor ^= x
		count[xor]++
	}
	return int64(res)
}
