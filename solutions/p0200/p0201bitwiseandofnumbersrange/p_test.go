package p0201bitwiseandofnumbersrange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	for _, tc := range []struct {
		left  int
		right int
		want  int
	}{
		{4, 5, 4},
		{2, 2, 2},
		{1, 2147483647, 0},
		{5, 7, 4},
		{0, 0, 0},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.left, tc.right), func(t *testing.T) {
			require.Equal(t, tc.want, rangeBitwiseAnd(tc.left, tc.right))
		})
	}
}

func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return right
	}
	res := 0
	var i int
	for i = 0; left > 0 && right > 0; i++ {
		if right&1 == 1 && right-left == 0 {
			res += 1 << i
		}
		right >>= 1
		left >>= 1
	}
	return res
}
