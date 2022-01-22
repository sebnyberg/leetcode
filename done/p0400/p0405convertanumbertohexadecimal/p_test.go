package p0405convertanumbertohexadecimal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_toHex(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want string
	}{
		{-1, "ffffffff"},
		{26, "1a"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, toHex(tc.num))
		})
	}
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	num32 := uint32(num)
	res := make([]byte, 0, 4)
	for num32 > 0 {
		r := num32 % 16
		if r <= 9 {
			res = append(res, byte(r+'0'))
		} else {
			res = append(res, byte(r-10+'a'))
		}
		num32 /= 16
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}
