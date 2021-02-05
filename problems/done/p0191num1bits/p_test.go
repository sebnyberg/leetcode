package p0191num1bits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hammingWeight(t *testing.T) {
	for _, tc := range []struct {
		num  uint32
		want int
	}{
		{3, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, hammingWeight(tc.num))
		})
	}
}

func hammingWeight(num uint32) (res int) {
	for num > 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}
