package p2571minimumoperationstoreduceanintegerto0

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{39, 3},
		{54, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.n))
		})
	}
}

func minOperations(n int) int {
	var count int
	for n > 0 {
		if n&1 == 0 {
			n >>= 1
			continue
		}
		if n&3 == 3 {
			n += 1
		}
		n >>= 1
		count++
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
