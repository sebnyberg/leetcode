package p0869reorderedpowerof2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reorderedPowerOf2(t *testing.T) {
	for _, tc := range []struct {
		N    int
		want bool
	}{
		{1, true},
		{10, false},
		{16, true},
		{24, false},
		{46, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.N), func(t *testing.T) {
			require.Equal(t, tc.want, reorderedPowerOf2(tc.N))
		})
	}
}

func reorderedPowerOf2(n int) bool {
	var numCount [10]int
	for x := n; x > 0; x /= 10 {
		numCount[x%10]++
	}
	for a := 1; a <= (1 << 32); a <<= 1 {
		var wantCount [10]int
		for x := a; x > 0; x /= 10 {
			wantCount[x%10]++
		}
		if wantCount == numCount {
			return true
		}
	}
	return false
}
