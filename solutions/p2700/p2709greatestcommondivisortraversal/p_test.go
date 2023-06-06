package p2709greatestcommondivisortraversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canTraverseAllPairs(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{3, 9, 5}, false},
		{[]int{2, 3, 6}, true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, canTraverseAllPairs(tc.nums))
		})
	}
}

func canTraverseAllPairs(nums []int) bool {
	// A non-1 gcd is a shared prime. For each number, factorize its primes and
	// gcd the primes together.
	// sort.Ints(nums)
	var factors []int
	factorize := func(x int) []int {
		factors = factors[:0]
		for y := 2; y*y <= x; y++ {
			if x%y != 0 {
				continue
			}
			factors = append(factors, y)
			for x%y == 0 {
				x /= y
			}
		}
		if x != 1 {
			factors = append(factors, x)
		}
		return factors
	}
	parent := make([]int, 1e5+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] == a {
			return a
		}
		ra := find(parent[a])
		parent[a] = ra
		return ra
	}

	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		parent[rb] = ra
	}

	seenPrimes := make(map[int]bool)
	for _, x := range nums {
		if x == 1 {
			if len(nums) > 1 {
				return false
			}
			return true
		}
		ds := factorize(x)
		for i := 0; i < len(ds); i++ {
			seenPrimes[ds[i]] = true
			for j := i + 1; j < len(ds); j++ {
				union(ds[i], ds[j])
			}
		}
	}
	root := -1
	for p := range seenPrimes {
		if root == -1 {
			root = find(p)
		} else {
			if find(p) != root {
				return false
			}
		}
	}
	return true
}
