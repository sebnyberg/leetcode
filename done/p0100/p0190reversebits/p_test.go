package p0190reversebits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseBits(t *testing.T) {
	for _, tc := range []struct {
		num  uint32
		want uint32
	}{
		{43261596, 964176192},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, reverseBits(tc.num))
		})
	}
}

func reverseBits(num uint32) uint32 {
	var reversed uint32
	for i := 0; i < 32; i++ {
		reversed <<= 1
		reversed += num % 2
		num >>= 1
	}
	return reversed
}
