package p2122recovertheoriginalarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_recoverArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{11, 6, 3, 4, 8, 7, 8, 7, 9, 8, 9, 10, 10, 2, 1, 9}, []int{2, 3, 7, 8, 8, 9, 9, 10}},
		{[]int{2, 10, 6, 4, 8, 12}, []int{3, 7, 11}},
		{[]int{1, 1, 3, 3}, []int{2, 2}},
		{[]int{5, 435}, []int{220}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, recoverArray(tc.nums))
		})
	}
}

func recoverArray(nums []int) []int {
	// We know:
	// The largest element in nums has to be part of arr[i]+k
	// The smallest element in nums has to be part of arr[i]-k
	// If a valid K exists, then it must be the difference between the smallest
	// element and some other element.
	numCount := make(map[int]int)
	recount := func() {
		for k := range numCount {
			numCount[k] = 0
		}
		for _, num := range nums {
			numCount[num]++
		}
	}
	sort.Ints(nums)
	res := make([]int, 0)
	// Never has to go to len(nums), but whatevss
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[0] || (nums[i]-nums[0])%2 != 0 {
			continue
		}
		twoK := nums[i] - nums[0]
		recount()
		for j := 0; j < len(nums); j++ {
			if numCount[nums[j]] > 0 {
				numCount[nums[j]+twoK] -= numCount[nums[j]]
				numCount[nums[j]] = 0
				if numCount[nums[j]+twoK] < 0 {
					goto continueSearch
				}
			}
		}
		recount()
		for j := 0; j < len(nums); j++ {
			if numCount[nums[j]] > 0 {
				numCount[nums[j]+twoK]--
				numCount[nums[j]]--
				res = append(res, nums[j]+(twoK/2))
			}
		}
		return res
	continueSearch:
	}
	return []int{}
}
