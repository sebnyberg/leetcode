package p0503base7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convertToBase7(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want string
	}{
		{100, "202"},
		{-7, "-10"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, convertToBase7(tc.num))
		})
	}
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	sign := 1
	if num < 0 {
		sign = -1
		num = -num
	}
	var res []byte
	for num > 0 {
		res = append(res, byte(num%7+'0'))
		num /= 7
	}
	if sign == -1 {
		res = append(res, '-')
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return string(res)
}
