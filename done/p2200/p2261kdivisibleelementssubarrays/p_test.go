package p2261kdivisibleelementssubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countDistinct(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k, p int
		want int
	}{
		{[]int{16, 17, 4, 12, 3}, 4, 1, 14},
		{[]int{2, 3, 3, 2, 2}, 2, 2, 11},
		{[]int{1, 2, 3, 4}, 4, 1, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countDistinct(tc.nums, tc.k, tc.p))
		})
	}
}

type trieNode struct {
	next [201]*trieNode
}

func countDistinct(nums []int, k int, p int) int {
	// There are max 4000 distinct subarrays in total
	// So we can just try all possible subarrays, just keeping track of whether
	// the sum of divisible elements is legal
	root := &trieNode{}
	var res int
	for i := 0; i < len(nums); i++ {
		var divCount int
		curr := root
		for j := i; j < len(nums); j++ {
			if nums[j]%p == 0 {
				divCount++
			}
			if divCount > k {
				goto ContinueSearch
			}
			if curr.next[nums[j]] == nil {
				curr.next[nums[j]] = &trieNode{}
				res++
			}
			curr = curr.next[nums[j]]
		}
	ContinueSearch:
	}
	return res
}
