package p2488countsubarrayswithmediank

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubarrays(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 5, 1, 4, 3, 6}, 1, 3},
		{[]int{3, 2, 1, 4, 5}, 4, 3},
		{[]int{2, 3, 1}, 3, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countSubarrays(tc.nums, tc.k))
		})
	}
}

func countSubarrays(nums []int, k int) int {
	// Funny leetcode does not follow the usual median definition. It made me
	// waste time.
	//
	// We want to calculate possible deltas on both sides for sequences of odd
	// and even length
	//
	evenCount := make(map[int]int)
	oddCount := make(map[int]int)

	var pos int
	for i := range nums {
		if nums[i] == k {
			pos = i
			break
		}
	}

	// Count on right side
	var delta int
	for i := pos + 1; i < len(nums); i++ {
		n := i - (pos + 1) + 1
		if nums[i] > k {
			delta++
		} else {
			delta--
		}
		if n&1 == 1 {
			oddCount[delta]++
		} else {
			evenCount[delta]++
		}
	}
	evenCount[0]++ // can always pick nothing
	res := evenCount[0]
	// We can also add odd counts for which there is one extra number above
	res += oddCount[1]

	// Count on left side
	delta = 0
	for i := pos - 1; i >= 0; i-- {
		n := (pos - 1) - i + 1
		if nums[i] > k {
			delta++
		} else {
			delta--
		}
		if n&1 == 1 {
			// There's an odd number of elements on the left side
			// We can pair that with any even count on the right for which the
			// combination yields -delta+1
			a := evenCount[-delta+1]
			res += a
			// And with any odd count that yields 0
			a = oddCount[-delta]
			res += a
		} else {
			// Even number of elements on left
			// Can combine with odd count that yields -delta+1
			a := oddCount[-delta+1]
			res += a

			// And even count that yields 0
			a = evenCount[-delta]
			res += a
		}
	}
	return res
}
