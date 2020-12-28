package d28_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reachNumber(t *testing.T) {
	for _, tc := range []struct {
		in   int
		want int
	}{
		// {3, 2},
		{2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, reachNumber(tc.in))
		})
	}
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	var n, sum int
	for sum < target {
		n++
		sum += n
	}
	delta := sum - target
	if delta&1 == 0 {
		return n
	}
	if n&1 == 0 {
		return n + 1
	}
	return n + 2
}
