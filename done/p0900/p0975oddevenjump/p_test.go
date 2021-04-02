package p0975oddevenjump

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_oddEvenJumps(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, oddEvenJumps(tc.arr))
		})
	}
}

func oddEvenJumps(arr []int) int {
	return 0
}
