package uber1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pancakeSort(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want []int
	}{
		{[]int{3, 2, 4, 1}, []int{3, 4, 2, 3, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, pancakeSort(tc.arr))
		})
	}
}

func pancakeSort(arr []int) []int {
	flips := make([]int, 0)
	r := len(arr) - 1
	for r > 0 {
		var maxIdx, maxVal int
		// Find max val
		for i := 0; i <= r; i++ {
			if arr[i] > maxVal {
				maxVal = arr[i]
				maxIdx = i
			}
		}
		// Flip max index into first position
		if maxIdx != 0 {
			flips = append(flips, maxIdx+1)
			rev(arr[:maxIdx+1])
		}
		// Flip max index into last position
		rev(arr[:r+1])
		flips = append(flips, r+1)
		r--
	}
	return flips
}

func rev(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}

func Test_findUnsortedSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		// {[]int{1, 3, 4, 2, 5}, 3},
		// {[]int{1, 2, 4, 5, 3}, 3},
		// {[]int{2, 3, 3, 2, 4}, 3},
		// {[]int{1, 3, 2, 2, 2}, 4},
		// {[]int{1, 2, 3, 3, 3}, 0},
		// {[]int{1, 2, 3, 4}, 0},
		// {[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		// {[]int{1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findUnsortedSubarray(tc.nums))
		})
	}
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for ; l < r && nums[l] <= nums[l+1]; l++ {
	}
	var newL int
	for newL := 0; nums[newL] <= nums[l]; newL++ {
		break
	}
	l = newL

	for ; l < r && nums[r] >= nums[l] && nums[r] >= nums[r-1]; r-- {
	}
	var newR int
	for newR := len(nums) - 1; nums[newR] >= nums[r]; newR-- {
		break
	}
	r = newR

	if r == l {
		return 0
	}
	for ; r < len(nums)-1 && nums[r] == nums[r+1]; r++ {
	}
	for ; l > 0 && nums[l] == nums[l-1]; l-- {
	}
	return r - l + 1
}
