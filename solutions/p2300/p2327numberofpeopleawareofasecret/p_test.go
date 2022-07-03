package p2327numberofpeopleawareofasecret

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_peopleAwareOfSecret(t *testing.T) {
	for _, tc := range []struct {
		n      int
		delay  int
		forget int
		want   int
	}{
		{6, 2, 4, 5},
		{4, 1, 3, 6},
		{425, 81, 118, 1754995},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, peopleAwareOfSecret(tc.n, tc.delay, tc.forget))
		})
	}
}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	// Keep track of the delta, and the total number of people
	// The total number of people is calculated by adding the delta each round
	if n < delay {
		return 0
	}
	const mod = 1e9 + 7
	var dp [1001]int
	people := 1
	for i := delay; i < forget; i++ {
		dp[i] = 1
	}
	dp[forget] -= 1
	for i := delay; i < n; i++ {
		for j := i + delay; j <= n && j <= i+forget; j++ {
			dp[j] = (dp[j] + dp[i]) % mod
		}
		if i+forget < n {
			dp[i+forget] -= dp[i]
		}
		people = (people + dp[i]) % mod
	}
	return people
}
