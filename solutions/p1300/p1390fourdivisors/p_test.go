package p1390fourdivisors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumFourDivisors(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{90351, 90779, 36358, 75474, 32986}, 147258},
		{[]int{90779, 36358, 90351, 75474, 32986}, 147258},
		{[]int{21, 4, 7}, 32},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, sumFourDivisors(tc.nums))
		})
	}
}

func sumFourDivisors(nums []int) int {
	// Finding all divisors for a number is the same as finding the integer
	// composition of a number. This is a pretty hard problem to do - it
	// involves finding all valid combinations of the prime factors of the
	// number.
	var res int
	for _, x := range nums {
		f := factors(x)
		if len(f) == 4 {
			for y := range f {
				res += y
			}
		}
	}
	return res
}

func factors(x int) map[int]struct{} {
	m := map[int]struct{}{
		1: {},
		x: {},
	}
	next := map[int]struct{}{}
	a := x
	for y := 2; y*y <= a; y++ {
		if a%y != 0 {
			continue
		}
		a /= y
		for k := range next {
			delete(next, k)
		}
		for k := range m {
			next[k] = struct{}{}
			if k <= y {
				if k*y <= x {
					next[k*y] = struct{}{}
				}
			} else if k%y == 0 {
				next[k/y] = struct{}{}
			}
		}
		m, next = next, m
		y--
	}
	return m
}
