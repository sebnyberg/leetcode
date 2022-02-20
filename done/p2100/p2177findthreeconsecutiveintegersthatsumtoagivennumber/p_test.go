package p2177

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumOfThree(t *testing.T) {
	for _, tc := range []struct {
		num  int64
		want []int64
	}{
		{33, []int64{10, 11, 12}},
		{4, []int64{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, sumOfThree(tc.num))
		})
	}
}

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	mid := num / 3
	return []int64{mid - 1, mid, mid + 1}
}
