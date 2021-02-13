package p0027removeelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 3}, 3, 0},
		// {[]int{3, 2, 2, 3}, 3, 2},
		// {[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.nums, tc.val), func(t *testing.T) {
			require.Equal(t, tc.want, removeElement(tc.nums, tc.val))
		})
	}
}

// Remove all instances of val in-place and return the new length
func removeElement(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			n--
			if i < n {
				nums[i] = nums[n]
				continue
			}
		}
		i++
	}
	// nums = nums[:i+1]
	return n
}
