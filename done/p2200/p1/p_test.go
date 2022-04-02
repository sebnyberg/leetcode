package p1

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minBitFlips(t *testing.T) {
	for _, tc := range []struct {
		start int
		goal  int
		want  int
	}{
		{10, 7, 3},
		{3, 4, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.start), func(t *testing.T) {
			require.Equal(t, tc.want, minBitFlips(tc.start, tc.goal))
		})
	}
}

func minBitFlips(start int, goal int) int {
	a := start ^ goal
	return bits.OnesCount(uint(a))
}
