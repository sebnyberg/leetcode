package p1611minimumonebitoperationstomakeintegerszero

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumOneBitOperations(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{8, 15},
		{12, 8},
		{326, 388},
		{6, 4},
		{3, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumOneBitOperations(tc.n))
		})
	}
}

func minimumOneBitOperations(n int) int {
	// Parse bits, MSB first
	bits := []int{}
	for x := n; x > 0; x >>= 1 {
		bits = append(bits, x&1)
	}
	for l, r := 0, len(bits)-1; l < r; l, r = l+1, r-1 {
		bits[l], bits[r] = bits[r], bits[l]
	}
	if n == 0 {
		return 0
	}
	costs := make([]int, 0)
	costs = append(costs, 0, 1)
	for i := 1; i < 32; i++ {
		costs = append(costs, costs[i]*2+1)
	}

	res := zero(costs, bits)

	return res
}

// returns cost to make bits into 0000...
func zero(costs, bits []int) int {
	for len(bits) > 1 && bits[0] == 0 {
		bits = bits[1:]
	}
	if len(bits) == 1 {
		return bits[0]
	}
	// Bits start with 1
	// Cost is equivalent to getting to 110... -> 010... -> 000...
	oneCost := one(costs, bits[1:])
	zeroCost := costs[len(bits)-1]
	return 1 + oneCost + zeroCost
}

// returns cost to make bits into 10000...
func one(costs, bits []int) int {
	if len(bits) == 1 {
		return 1 - bits[0]
	}
	if bits[0] == 0 {
		oneCost := one(costs, bits[1:])
		zeroCost := costs[len(bits)-1]
		return 1 + oneCost + zeroCost
	}
	return zero(costs, bits[1:])
}
