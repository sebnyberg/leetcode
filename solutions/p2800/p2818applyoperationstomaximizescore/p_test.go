package p2818applyoperationstomaximizescore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumScore(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{8, 3, 9, 3, 8}, 2, 81},
		{[]int{19, 12, 14, 6, 10, 18}, 3, 4788},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumScore(tc.nums, tc.k))
		})
	}
}

func maximumScore(nums []int, k int) int {
	// There are two subproblems:
	//
	// 1. Calculate the prime score of each element in nums
	// 2. Select subarrays in an optimal way to maximise the total score
	//
	// Let's start with prime factorization
	numPrimes := make(map[int]int)
	for _, x := range nums {
		if _, exists := numPrimes[x]; exists {
			continue
		}
		for factor := 2; factor*factor <= x; factor++ {
			if x%factor != 0 {
				continue
			}
			numPrimes[x]++
			for x%factor == 0 {
				x /= factor
			}
		}
	}

	// Second subproblem is to figure out which subarrays to select
}
