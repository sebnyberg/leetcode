package p0401binarywatch

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_readBinaryWatch(t *testing.T) {
	for _, tc := range []struct {
		turnedOn int
		want     []string
	}{
		{1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
		{9, []string{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.turnedOn), func(t *testing.T) {
			require.Equal(t, tc.want, readBinaryWatch(tc.turnedOn))
		})
	}
}

func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)
	for min := uint8(0); min <= 11; min++ {
		for second := uint8(0); second <= 59; second++ {
			if bits.OnesCount8(min)+bits.OnesCount8(second) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", min, second))
			}
		}
	}
	return res
}
