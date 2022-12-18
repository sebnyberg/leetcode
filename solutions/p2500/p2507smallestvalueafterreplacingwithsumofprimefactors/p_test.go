package p2507smallestvalueafterreplacingwithsumofprimefactors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestValue(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{4, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, smallestValue(tc.n))
		})
	}
}

func smallestValue(n int) int {
	findFactor := func(x int) int {
		for a := 2; a*a <= x; a++ {
			if x%a == 0 {
				return a
			}
		}
		return x
	}
	prev := n
	curr := n
	for {
		fac := findFactor(curr)
		if fac == curr {
			break
		}
		var sum int
		for {
			fac := findFactor(curr)
			sum += fac
			if fac == curr {
				break
			}
			curr /= fac
		}
		curr = sum
		if curr == prev {
			break
		}
		prev = curr
	}
	return curr
}
