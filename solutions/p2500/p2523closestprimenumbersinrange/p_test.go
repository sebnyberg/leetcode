package p4

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closestPrimes(t *testing.T) {
	for i, tc := range []struct {
		left  int
		right int
		want  []int
	}{
		{10, 19, []int{11, 19}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, closestPrimes(tc.left, tc.right))
		})
	}
}

func closestPrimes(left int, right int) []int {
	// Sieve of Eratosthenes
	notPrime := make([]bool, right+1)
	var prev int
	minDist := math.MaxInt32
	res := []int{-1, -1}
	for x := 2; x <= right; x++ {
		if notPrime[x] {
			continue
		}
		for y := x; y <= right; y += x {
			notPrime[y] = true
		}
		if x < left {
			continue
		}
		if prev != 0 {
			d := x - prev
			if d < minDist {
				minDist = d
				res = []int{prev, x}
			}
		}
		prev = x
	}
	return res
}
