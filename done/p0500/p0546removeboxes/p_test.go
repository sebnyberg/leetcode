package p0546removeboxes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeBoxes(t *testing.T) {
	for _, tc := range []struct {
		boxes []int
		want  int
	}{
		{[]int{1, 3, 2, 2, 2, 3, 4, 3, 1}, 23},
		{[]int{1, 1, 1}, 9},
		{[]int{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.boxes), func(t *testing.T) {
			require.Equal(t, tc.want, removeBoxes(tc.boxes))
		})
	}
}

func removeBoxes(boxes []int) int {

}
