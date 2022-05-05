package p0372superpow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_superPow(t *testing.T) {
	for _, tc := range []struct {
		a    int
		b    []int
		want int
	}{
		{2147483647, []int{2, 0, 0}, 1198},
		{2, []int{3}, 8},
		{2, []int{1, 0}, 1024},
		{1, []int{4, 3, 3, 8, 5, 2}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.a), func(t *testing.T) {
			require.Equal(t, tc.want, superPow(tc.a, tc.b))
		})
	}
}

const mod = 1337

func superPow(a int, b []int) int {
	if len(b) == 1 && b[0] == 0 {
		return 1
	}
	// Calculate a pow b mod 1337 where b is an extremely large positive integer
	// in the form of an array.
	// Idea: iteratively divide b by 2 and check rest. This gives an array in
	// the form [1,0,1,1,0,0,0,1]. What it means in practice is that dividing the
	// entire exponent by half resulted in an extra operation.
	// Then, iterate over the rests array from last to first element, performing
	// bottom-up modPow.
	var rest int
	rests := make([]int, 0)
	for len(b) > 0 {
		b, rest = div2Rest(b)
		rests = append(rests, rest)
	}

	base := a % mod
	for i := len(rests) - 2; i >= 0; i-- {
		res := base * base % mod
		if rests[i] == 1 {
			res = res * a % mod
		}
		base = res
	}
	return base
}

func div2Rest(b []int) ([]int, int) {
	var carry int
	for i := 0; i < len(b); i++ {
		v := b[i] + carry*10
		carry = v % 2
		b[i] = v / 2
	}
	if b[0] == 0 {
		b = b[1:]
	}
	return b, carry
}
