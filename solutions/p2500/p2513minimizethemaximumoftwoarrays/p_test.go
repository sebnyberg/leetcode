package p2513minimizethemaximumoftwoarrays

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimizeSet(t *testing.T) {
	for i, tc := range []struct {
		divisor1   int
		divisor2   int
		uniqueCnt1 int
		uniqueCnt2 int
		want       int
	}{
		{2, 7, 1, 3, 4},
		{3, 5, 2, 1, 3},
		{2, 4, 8, 2, 15},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimizeSet(tc.divisor1, tc.divisor2, tc.uniqueCnt1, tc.uniqueCnt2))
		})
	}
}

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	// Let's try binary search
	//
	// Given a claimed solution, we can check whether the solution is possible
	// or not.
	//
	// Any number that is % divisor1 should be assigned to the second set,
	// unless also % divisor2. The number that is % divisor1 and divisor2 is the
	// LCM.
	//
	// And vice versa.
	//
	// Any other number must be assignable to either. So the remaining sum of
	// numbers missing from both sets must be <= the remaining numbers.

	lcm := divisor1 * divisor2 / gcd(divisor1, divisor2)
	check := func(x int) bool {
		// Not assignable to either sets
		notAssignable := x / lcm

		// Must be assigned to first set
		assignFirst := (x / divisor2) - notAssignable
		assignSecond := (x / divisor1) - notAssignable

		cnt1 := max(0, uniqueCnt1-assignFirst)
		cnt2 := max(0, uniqueCnt2-assignSecond)

		assignEither := x - notAssignable - assignFirst - assignSecond
		return assignEither >= cnt1+cnt2
	}

	lo := 0
	hi := math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
