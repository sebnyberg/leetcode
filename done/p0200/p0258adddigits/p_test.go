package p0258adddigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addDigits(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{38, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, addDigits(tc.num))
		})
	}
}

func addDigits(num int) int {
	for num >= 10 {
		res := 0
		for num >= 10 {
			res *= 10
			res += num % 10
			num /= 10
		}
		num = num + res
	}
	return num
}
