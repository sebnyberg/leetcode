package p231poweroftwo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfSteps(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want bool
	}{
		{0, false},
		{1, true},
		{16, true},
		{3, false},
		{4, true},
		{5, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, isPowerOfTwo(tc.num))
		})
	}
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&-n == n
}
