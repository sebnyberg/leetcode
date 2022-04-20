package p0717p1bitand2bitcharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isOneBitCharacter(t *testing.T) {
	for _, tc := range []struct {
		bits []int
		want bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bits), func(t *testing.T) {
			require.Equal(t, tc.want, isOneBitCharacter(tc.bits))
		})
	}
}

func isOneBitCharacter(bits []int) bool {
	var oneBit bool
	for i := 0; i < len(bits); {
		if bits[i] == 1 {
			i += 2
			oneBit = false
		} else {
			i++
			oneBit = true
		}
	}
	return oneBit
}
