package p1893a

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_chalkReplacer(t *testing.T) {
	for _, tc := range []struct {
		chalk []int
		k     int
		want  int
	}{
		{[]int{5, 1, 5}, 22, 0},
		{[]int{3, 4, 1, 2}, 25, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.chalk), func(t *testing.T) {
			require.Equal(t, tc.want, chalkReplacer(tc.chalk, tc.k))
		})
	}
}

func chalkReplacer(chalk []int, k int) int {
	idx := 0
	n := len(chalk)
	oneRound := 0
	for _, c := range chalk {
		oneRound += c
	}
	k %= oneRound
	for {
		k -= chalk[idx]
		if k < 0 {
			break
		}
		idx++
		idx %= n
	}
	return idx
}
