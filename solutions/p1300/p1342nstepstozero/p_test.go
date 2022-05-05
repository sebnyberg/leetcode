package p1342nstepstozero

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfSteps(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{14, 6},
		{8, 4},
		{123, 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfSteps(tc.num))
		})
	}
}

func numberOfSteps(num int) (actions int) {
	for num > 0 {
		if num&1 == 1 {
			num--
		} else {
			num /= 2
		}
		actions++
	}
	return actions
}
