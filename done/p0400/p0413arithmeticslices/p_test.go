package p0413arithmeticslices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfArithmeticSlices(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		want int
	}{
		{[]int{1, 3, 5, 7, 9}, 6},
		{[]int{7, 7, 7, 7}, 3},
		{[]int{3, -1, -5, -9}, 3},
		{[]int{1, 2, 3, 4}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfArithmeticSlices(tc.A))
		})
	}
}

func numberOfArithmeticSlices(A []int) (res int) {
	if len(A) <= 1 {
		return res
	}
	runLength := 2
	diff := A[1] - A[0]
	for i := 2; i < len(A); i++ {
		d := A[i] - A[i-1]
		if d != diff {
			diff = d
			runLength = 2
			continue
		}
		runLength++
		res += runLength - 2
	}
	return res
}
