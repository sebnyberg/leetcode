package p1551minoptomakearrayequal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{1, 0},
		{3, 2},
		{4, 4},
		{6, 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.n))
		})
	}
}

// This can be made more efficient by using the geometric sum of
// i=start -> n-1 for (2*i - n + 1)
// Which becomes
// i=0 -> n-1 for (2*i - n + 1) - i=start -> n-1 for (2*i - n + 1)
// Given i=0->m for k => m(m-1)/2, insert the above and get immediate results.
// I just cba to do the calculation and this was 0ms so left it there.
func minOperations(n int) int {
	var sum int
	start := n/2 + n%2
	for i := start; i < n; i++ {
		sum += 2*i + 1 - n
	}
	return sum
}
