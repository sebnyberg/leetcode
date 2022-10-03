package p2429maximizexor

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimizeXor(t *testing.T) {
	for i, tc := range []struct {
		num1 int
		num2 int
		want int
	}{
		{1, 12, 3},
		{3, 5, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimizeXor(tc.num1, tc.num2))
		})
	}
}

func minimizeXor(num1 int, num2 int) int {
	// nbits1 := bits.OnesCount(uint(num2))
	nbits2 := bits.OnesCount(uint(num2))

	a := []byte(fmt.Sprintf("%b", num1))

	// There are a couple of cases here
	b := make([]byte, max(len(a), nbits2))
	for i := range b {
		b[i] = '0'
	}

	// Let's fill resBytes with 1s wherever we should put ones.
	//
	// First, let's prioritize removing 1s from num1
	for i := range a {
		if a[i] == '1' && nbits2 > 0 {
			b[i] = '1' // 'erase' the 1 from nums1
			nbits2--
		}
	}
	// Then, fill any remaining zeroes from low to high until nbits == 0
	for i := len(b) - 1; i >= 0; i-- {
		if nbits2 > 0 && b[i] == '0' {
			b[i] = '1'
			nbits2--
		}
	}
	var res int
	for i := 0; i < len(b); i++ {
		res <<= 1
		res += int(b[i] - '0')
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
