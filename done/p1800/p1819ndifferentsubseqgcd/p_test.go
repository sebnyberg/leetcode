package p1819ndifferentsubseqgcd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countDifferentSubsequenceGCDs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{6, 10, 3}, 5},
		{[]int{5, 15, 40, 5, 6}, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countDifferentSubsequenceGCDs(tc.nums))
		})
	}
}

func countDifferentSubsequenceGCDs(nums []int) int {
	// maxnum := int(math.Sqrt(1e5))
	n := len(nums)
	ns := make(map[int]bool, n)
	for _, n := range nums {
		ns[n] = true
	}
	res := 0
	// For each possible GCD (i)
	for i := 1; i <= 200001; i++ {
		cur := 0
		// Check if there is a number k which gcd with another number from
		// our series becomes i. Then i is a GCD for a subsequence.
		for k := i; k <= 200001; k += i {
			if !ns[k] {
				continue
			}
			if cur == 0 {
				cur = k
			} else {
				cur = gcd(cur, k)
			}
			if cur == i {
				res++
				break
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
