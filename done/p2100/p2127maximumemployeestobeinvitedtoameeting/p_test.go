package p2127maximumemployeestobeinvitedtoameeting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumInvitations(t *testing.T) {
	for _, tc := range []struct {
		favorite []int
		want     int
	}{
		{[]int{2, 2, 1, 2}, 3},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 0, 1, 4, 1}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.favorite), func(t *testing.T) {
			require.Equal(t, tc.want, maximumInvitations(tc.favorite))
		})
	}
}

func maximumInvitations(favorite []int) int {

}
