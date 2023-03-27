package p2601primesubtractionoperation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_primeSubOperation(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{998, 2}, true},
		{[]int{4, 9, 6, 10}, true},
		{[]int{6, 8, 11, 12}, true},
		{[]int{5, 8, 3}, false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, primeSubOperation(tc.nums))
		})
	}
}

func primeSubOperation(nums []int) bool {
	// There aren't that many primes <= 1000, find them
	primes := []int{}
	var notprime [1001]bool
	for x := 2; x <= 1000; x++ {
		if notprime[x] {
			continue
		}
		notprime[x] = true
		primes = append(primes, x)
		for y := x * 2; y <= 1000; y += x {
			notprime[y] = true
		}
	}

	// Remove largest possible prime from each value such that its value >=
	// previous value.
	for i := range nums {
		prev := 0
		if i > 0 {
			prev = nums[i-1]
		}
		p := -1
		for _, x := range primes {
			if nums[i]-x <= prev {
				continue
			}
			if x >= nums[i] {
				break
			}
			p = x
		}
		if p == -1 {
			if nums[i] > prev {
				continue
			}
			return false
		}
		nums[i] -= p
		prev = nums[i]
	}
	return true
}
