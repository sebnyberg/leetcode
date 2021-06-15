package p0473matchstickstosquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makesquare(t *testing.T) {
	for _, tc := range []struct {
		matchsticks []int
		want        bool
	}{
		{[]int{1, 1, 2, 2, 2}, true},
		{[]int{3, 3, 3, 3, 4}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matchsticks), func(t *testing.T) {
			require.Equal(t, tc.want, makesquare(tc.matchsticks))
		})
	}
}

func makesquare(matchsticks []int) bool {
	// Trying all options is 2^15 so that's OK
	sides := make([]int, 4)
	var sum int
	for _, stick := range matchsticks {
		sum += stick
	}
	if sum%4 != 0 {
		return false
	}
	return tryAll(matchsticks, sides, sum/4, 0, len(matchsticks))
}

func tryAll(matchSticks, sides []int, quarter, idx, n int) bool {
	if idx == n {
		return true
	}

	for i := 0; i < 4; i++ {
		if sides[i]+matchSticks[idx] > quarter {
			continue
		}
		sides[i] += matchSticks[idx]
		if tryAll(matchSticks, sides, quarter, idx+1, n) {
			return true
		}
		sides[i] -= matchSticks[idx]
	}

	return false
}
