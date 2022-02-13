package p2172maximumandsumofarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumANDSum(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		numSlots int
		want     int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, 9},
		{[]int{1, 3, 10, 4, 7, 1}, 9, 24},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumANDSum(tc.nums, tc.numSlots))
		})
	}
}

func maximumANDSum(nums []int, numSlots int) int {
	// Each slot can be empty, have one entry, two entries => base 3
	mask := 1
	for i := 0; i < numSlots; i++ {
		mask *= 3
	}
	mask -= 1
	mem := make([]int, mask+1)
	res := dp(mem, nums, len(nums)-1, mask, numSlots)
	return res
}

func dp(mem, nums []int, i, mask, numSlots int) int {
	if mem[mask] > 0 {
		return mem[mask]
	}
	if i < 0 {
		return 0
	}
	for slot, bit := 1, 1; slot <= numSlots; slot, bit = slot+1, bit*3 {
		if (mask/bit)%3 > 0 { // Places left in the slot
			mem[mask] = max(mem[mask], nums[i]&slot+dp(mem, nums, i-1, mask-bit, numSlots))
		}
	}
	return mem[mask]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
