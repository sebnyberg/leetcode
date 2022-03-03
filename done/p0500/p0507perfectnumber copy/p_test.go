package p0507perfectnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkPerfectNumber(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want bool
	}{
		{1, false},
		{28, true},
		{7, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, checkPerfectNumber(tc.num))
		})
	}
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	// A divisor of num must be at most its square
	res := 1 // include the divisor 1
	for x := 2; x*x <= num; x++ {
		if num%x != 0 {
			continue
		}
		res += x
		if num/x != x {
			res += num / x
		}
	}
	return num == res
}
