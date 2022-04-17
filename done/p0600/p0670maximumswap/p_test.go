package p0670maximumswap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSwap(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{2736, 7236},
		{9973, 9973},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSwap(tc.num))
		})
	}
}

func maximumSwap(num int) int {
	s := []byte(fmt.Sprint(num))
	var maxIdx int
	var maxNum byte
	for i := 0; i < len(s)-1; i++ {
		maxIdx = -1
		maxNum = s[i] + 1
		for j := i + 1; j < len(s); j++ {
			if s[j] >= maxNum {
				maxNum = s[j]
				maxIdx = j
			}
		}
		if maxIdx != -1 {
			s[i], s[maxIdx] = s[maxIdx], s[i]
			var x int
			for i := range s {
				x *= 10
				x += int(s[i] - '0')
			}
			return x
		}
	}
	return num
}
