package p2023numberofpairsofstringswithconcatenationsequaltotarget

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numOfPairs(t *testing.T) {
	for _, tc := range []struct {
		nums   []string
		target string
		want   int
	}{
		{[]string{"777", "7", "77", "77"}, "7777", 4},
		{[]string{"123", "4", "12", "34"}, "1234", 2},
		{[]string{"1", "1", "1"}, "11", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numOfPairs(tc.nums, tc.target))
		})
	}
}

func numOfPairs(nums []string, target string) int {
	var res int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res++
			}
			if nums[j]+nums[i] == target {
				res++
			}
		}
	}
	return res
}
