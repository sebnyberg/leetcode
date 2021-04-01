package p0717_1bit2bitchar

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
		{[]int{1, 1, 1, 0}, false},
		{[]int{1, 0, 0}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bits), func(t *testing.T) {
			require.Equal(t, tc.want, isOneBitCharacter(tc.bits))
		})
	}
}

func isOneBitCharacter(bits []int) bool {
	var i int
	var onebit bool
	for i = 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
			onebit = false
			continue
		}
		onebit = true
	}
	return onebit
}
