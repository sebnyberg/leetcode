package p0507perfectnumber

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPerfectNumber(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want bool
	}{
		{28, true},
		{6, true},
		{496, true},
		{8128, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, checkPerfectNumber(tc.num))
		})
	}
}

func checkPerfectNumber(n int) bool {
	if n == 1 {
		return false
	}
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			sum += i + n/i
		}
	}
	return sum == n
}
