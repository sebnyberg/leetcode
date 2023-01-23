package p2543checkifpointisreachable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isReachable(t *testing.T) {
	for i, tc := range []struct {
		targetX int
		targetY int
		want    bool
	}{
		{6, 9, false},
		{4, 7, true},
		{3, 7, true},
		{757172937, 869964136, true},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, isReachable(tc.targetX, tc.targetY))
		})
	}
}

func isReachable(targetX int, targetY int) bool {
	// See
	// https://leetcode.com/problems/check-if-point-is-reachable/solutions/3091988/another-detailed-explanation/
	// Euclidian algorithm
	for targetY != 0 {
		targetX, targetY = targetY, targetX%targetY
	}

	// Validate that gcd is a power-of-two
	gcd := targetX
	for gcd%2 == 0 {
		gcd /= 2
	}
	return gcd == 1
}
