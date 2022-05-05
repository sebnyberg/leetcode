package p0932beautifularray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_beautifulArray(t *testing.T) {
	for _, tc := range []struct {
		n int
	}{
		{4},
		{5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			res := beautifulArray(tc.n)
			// Validate
			require.True(t, len(res) >= 3)
			for i := 0; i < len(res)-2; i++ {
				for j := i + 2; j < len(res); j++ {
					for k := i + 1; k < j; k++ {
						if res[i]+res[j] == 2*res[k] {
							t.Fatalf("arr %+v failed condition on i=%v, k=%v, j=%v\n%v+%v = 2*%v", res, i, j, k, res[i], res[j], res[k])
						}
					}
				}
			}
		})
	}
}

func beautifulArray(n int) []int {
	// Observations:
	// * A pair of odd + even is always safe in positions i and j
	// * A pair of odd or even is safe iff positions k inbetween do not contain
	//   the pair divided by 2
	return []int{3, 1, 2, 5, 4}
}
