package p2195appendkintegerswithminimalsum

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimalKSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 4, 25, 10, 25}, 2, 5},
		{[]int{5, 6}, 6, 25},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimalKSum(tc.nums, tc.k))
		})
	}
}

func minimalKSum(nums []int, k int) int64 {
	// Sort and de-duplicate numbers in nums
	sort.Ints(nums)
	var j int
	for i := range nums {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	nums = nums[:j+1]
	nums = append(nums, 0)
	copy(nums[1:], nums)
	nums[0] = 0
	nums = append(nums, math.MaxInt64)

	// Add gaps from nums until reaching k
	var sum int
	for i := 1; i < len(nums); i++ {
		a, b := nums[i-1], nums[i]
		c := min(k, b-a-1)
		sum += c * (c + 1) / 2
		sum += a * c
		k -= c
		if k == 0 {
			break
		}
	}

	return int64(sum)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
