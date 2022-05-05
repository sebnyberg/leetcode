package p0163missingranges

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMissingRanges(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		lower int
		upper int
		want  []string
	}{
		{[]int{0, 1, 3, 50, 75}, 0, 99, []string{"2", "4->49", "51->74", "76->99"}},
		{[]int{}, 1, 1, []string{"1"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMissingRanges(tc.nums, tc.lower, tc.upper))
		})
	}
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		return []string{getRange(lower, upper)}
	}

	if lower < nums[0] {
		nums = append(nums, 1)
		copy(nums[1:], nums)
		nums[0] = lower - 1
	}

	n := len(nums) - 1
	for nums[n] > upper {
		n--
	}
	nums = nums[:n+1]
	nums = append(nums, upper+1)
	res := make([]string, 0)

	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d > 1 {
			res = append(res, getRange(nums[i-1]+1, nums[i]-1))
		}
	}

	return res
}

func getRange(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	}
	return fmt.Sprintf("%v->%v", a, b)
}
