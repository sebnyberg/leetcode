package p0292nimgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canWinNim(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{4, false},
		{1, true},
		{2, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, canWinNim(tc.n))
		})
	}
}

func canWinNim(n int) bool {
	return n%4 != 0
}
