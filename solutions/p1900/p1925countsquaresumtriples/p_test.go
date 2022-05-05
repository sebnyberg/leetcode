package p1925countsquaresumtriples

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countTriples(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{5, 2},
		{10, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countTriples(tc.n))
		})
	}
}

func countTriples(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if i*i+j*j == k*k {
					count++
				}
			}
		}
	}
	return count
}
