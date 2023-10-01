package p1359countallvalidpickupanddeliveryoptions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countOrders(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{2, 6},
		{1, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countOrders(tc.n))
		})
	}
}

const mod = 1e9 + 7

func countOrders(n int) int {
	// Given n orders, each order consist of a pickup and delivery service.
	//
	// Count all valid pickup/delivery possible sequences such that delivery(i)
	// is always after pickup(i).
	//
	// The first big thing to note is that the order of the pickup and delivery
	// doesn't matter, only how many options we have and how many pickups are
	// "active" at a given point in time.
	//
	// That is, if there are 7 pickups, then we know that we have 7 choices for
	// delivery if we wish to do so.
	//
	// This tells us that we can simply use DP to calculate the number of ways
	// in which we would find ourselves in a given state, where the state is
	// given by the number of active pickups.
	//
	// The only "tricky" part is that we need to take care not to go beyond n in
	// terms of total pickups.
	//
	dp := make([]int, n+1)
	dp[0] = 1
	next := make([]int, n+1)
	reset := func(a []int) {
		for i := range a {
			a[i] = 0
		}
	}
	for t := 0; t < 2*n; t++ {
		reset(next)
		for deliveries := max(-n+t, 0); deliveries <= t/2; deliveries++ {
			// Assuming that we've made "deliveries" deliveries, then we must've
			// also made at least as many pickups.
			pickups := t - deliveries

			// Given that we've made "pickups" pickups and "deliveries"
			// deliveries, how many ways can we make the next pickup?
			// The answer is "n-pickups" because there are that many options
			// left.
			possiblePickups := n - pickups
			pickupDelta := pickups - deliveries
			if pickups < n {
				next[pickupDelta+1] = (next[pickupDelta+1] + possiblePickups*dp[pickupDelta]) % mod
			}

			if pickupDelta > 0 {
				next[pickupDelta-1] = (next[pickupDelta-1] + pickupDelta*dp[pickupDelta]) % mod
			}
		}
		dp, next = next, dp
	}

	return dp[0] % mod
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
