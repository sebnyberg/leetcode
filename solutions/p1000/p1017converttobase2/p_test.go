package p1017converttobase2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_baseNeg2(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want string
	}{
		{2, "110"},
		{3, "111"},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, baseNeg2(tc.n))
		})
	}
}

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var res string
	for n != 0 {
		res = fmt.Sprint(n&1) + res
		n = -(n >> 1)
	}
	return res
}
