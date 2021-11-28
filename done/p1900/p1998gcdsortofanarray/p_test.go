package p1998gcdsortofanarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_gcdSort(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{7, 21, 3}, true},
		{[]int{5, 2, 6, 2}, false},
		{[]int{10, 5, 9, 3, 15}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, gcdSort(tc.nums))
		})
	}
}

func gcdSort(nums []int) bool {
	// A pair of numbers can be swapped if they share a prime factor
	maxVal := nums[0]
	for _, n := range nums[1:] {
		maxVal = max(maxVal, n)
	}

	// Pre-calculate prime factors using sieve
	smallestPrime := make([]int, maxVal+1)
	for i := 2; i < len(smallestPrime); i++ {
		smallestPrime[i] = i
	}
	for i := 2; i*i <= maxVal; i++ {
		for j := i; j <= maxVal; j += i {
			if smallestPrime[j] == j {
				smallestPrime[j] = i
			}
		}
	}

	// Set up DSU
	parent := make([]int, maxVal+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			root := find(parent[a])
			parent[a] = root
		}
		return parent[a]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = a
		}
	}

	// For each number in nums, union the number to each of its prime factors
	for _, n := range nums {
		numCpy := n
		for numCpy > 1 {
			union(n, smallestPrime[numCpy])
			numCpy /= smallestPrime[numCpy]
		}
	}

	// Figure out where each number should be in order to be sorted
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	// Finally, check if each number could be swapped to its sorted position based
	// on its set
	for i, n := range nums {
		if find(sortedNums[i]) != find(n) {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
