package p2498frogjumpii

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxJump(t *testing.T) {
	for i, tc := range []struct {
		stones []int
		want   int
	}{
		{[]int{0, 2, 5, 6, 7}, 5},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxJump(tc.stones))
		})
	}
}

func maxJump(stones []int) int {
	// My intuition is telling me that given a maximum length jump, you should
	// always jump as far as possible in both directions such that the actual
	// jump <= maxJump
	//
	// So we could just test whether a cestain length is ok or not.
	n := len(stones)
	seen := make([]bool, n)
	ok := func(jumpLen int) bool {
		for i := range seen {
			seen[i] = false
		}
		// Go forward
		seen[0] = true
		var i int
		for i != n-1 {
			if stones[i+1]-stones[i] > jumpLen {
				return false
			}
			// Jump as far as possible
			j := i + 1
			for {
				if j == n-1 {
					break
				}
				if stones[j+1]-stones[i] > jumpLen {
					break
				}
				j++
			}
			seen[j] = true
			i = j
		}
		// Go back and do any possible jump
		for i != 0 {
			j := i - 1
			for j > 0 && seen[j] {
				j--
			}
			if stones[i]-stones[j] > jumpLen {
				return false
			}
			i = j
		}
		return true
	}
	lo, hi := 0, math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !ok(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
